package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"moshel-api/lib"
	"moshel-api/models"

	"github.com/gin-gonic/gin"
)

var secretKey = lib.GetEnv("SECRET_KEY")
var costFactor = lib.GetEnv("COST_FACTOR")

func Register(c *gin.Context) (string, error) {
	db := c.MustGet("db").(*sql.DB)
	var UserInput models.UserDataInput

	if err := c.ShouldBind(&UserInput); err != nil {
		return "", err
	}

	if UserInput.Password != UserInput.ConfirmPassword {
		errMsg := errors.New("Password is not match, please try again")
		return "", errMsg
	}

	encryptedPass, err := lib.HashPassword(UserInput.Password, costFactor)

	if UserInput.Password != UserInput.ConfirmPassword {
		errMsg := errors.New(err.Error())
		return "", errMsg
	}

	query := fmt.Sprintf("INSERT INTO users(username, password, created_at) values('%v', '%v', NOW());", UserInput.Username, encryptedPass)

	_, err = db.ExecContext(c, query)

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("User already registered")
	}

	token, err := lib.GenerateToken(UserInput.Username, secretKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(c *gin.Context) (string, error) {
	db := c.MustGet("db").(*sql.DB)
	var userInput models.UserDataInput

	if err := c.ShouldBind(&userInput); err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT username, password FROM users where username = '%v'", userInput.Username)
	var username, password string

	err := db.QueryRowContext(c, query).Scan(&username, &password)
	log.Println("username", username)
	log.Println("pass", password)
	if err != nil {
		log.Println("error", err.Error())
		msg := fmt.Sprintf("Cannot find user with username '%v'", userInput.Username)
		return "", errors.New(msg)
	}

	if err := lib.UnHashPassword(userInput.Password, password); err != nil {
		return "", err
	}

	token, err := lib.GenerateToken(username, secretKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

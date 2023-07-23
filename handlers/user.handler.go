package handlers

import (
	"database/sql"
	"errors"
	"fmt"
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

	getQuery := fmt.Sprintf("SELECT username FROM user WHERE username = ?")

	query := fmt.Sprintf("INSERT INTO user(username, password) value('%s', '%s');", UserInput.Username, encryptedPass)
	var resultUsername string

	db.QueryRow(getQuery, UserInput.Username).Scan(&resultUsername)

	if resultUsername == UserInput.Username {
		return "", errors.New("User already registered")
	}

	result , err := db.Query(query)

	defer result.Close()

	if err != nil {
		return "", errors.New("Cannot insert data to db")
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

	query := fmt.Sprintf("SELECT username, password FROM user where username = ?")
	var username, password string

	if err := db.QueryRow(query, userInput.Username).Scan(&username, &password); err != nil {
		return "", errors.New("User not registered")
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

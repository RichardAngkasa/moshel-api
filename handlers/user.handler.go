package handlers

import (
	"errors"
	"fmt"
	"moshel-api/lib"
	"moshel-api/models"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

var secretKey = lib.GetEnv("SECRET_KEY")

func Register(c *gin.Context) (string, error) {
	var UserInput UserInput

	if err := c.ShouldBind(&UserInput); err != nil {
		return "", err
	}

	if UserInput.Password != UserInput.ConfirmPassword {
		errMsg := errors.New("Password is not match, please try again")
		return "", errMsg
	}

	encryptedPass, err := lib.EncryptPass(UserInput.Password, secretKey)

	if UserInput.Password != UserInput.ConfirmPassword {
		errMsg := errors.New(err.Error())
		return "", errMsg
	}

	NewUser := models.UserData{
		Username: UserInput.Username,
		Password: encryptedPass,
	}

	fmt.Println(NewUser)

	return "", nil
}

func Login(c *gin.Context) (string, error) {
	var userInput UserInput

	if err := c.ShouldBind(&userInput); err != nil {
		return "", err
	}
	textPass, _ := lib.Decrypt(userInput.Password, secretKey)

	return textPass, nil
}

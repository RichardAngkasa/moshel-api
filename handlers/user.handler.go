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

func CreateUser(c *gin.Context) (string, error) {
	var UserInput UserInput

	if err := c.ShouldBind(&UserInput); err != nil {
		return "", err
	}

	if UserInput.Password != UserInput.ConfirmPassword {
		errMsg := errors.New("Password is not match, please try again")
		return "", errMsg
	}

	encryptedPass, err := lib.EncryptPass(UserInput.Password)

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

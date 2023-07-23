package controllers

import (
	"moshel-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	token, err := handlers.Register(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H { "message": "success", "token": token })
}

func GetUser(c *gin.Context) {
	token, err := handlers.Login(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusOK, gin.H { "message": "success", "token": token })
}

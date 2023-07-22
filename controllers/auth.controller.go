package controllers

import (
	"moshel-api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	token, err := handlers.CreateUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	c.JSON(http.StatusCreated, gin.H { "message": token })
}

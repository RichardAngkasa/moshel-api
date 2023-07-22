package router

import (
	"moshel-api/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/user")
	auth.POST("/register", controllers.CreateUser)

	return r
}

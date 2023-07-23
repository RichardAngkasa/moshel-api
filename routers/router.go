package router

import (
	"database/sql"
	"moshel-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(func (c *gin.Context) {
		c.Set("db", db)
	})

	auth := r.Group("/user")
	auth.POST("/register", controllers.CreateUser)
	auth.POST("/login", controllers.GetUser)

	return r
}

package main

import (
	"log"
	router "moshel-api/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	r := router.CreateRoutes()

	r.Run(":8080")
}

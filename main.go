package main

import (
	"log"
	"moshel-api/lib"
	router "moshel-api/routers"
)

func main() {
	db := lib.ConnectDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	r := router.CreateRoutes(db)

	r.Run()
}

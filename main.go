package main

import (
	router "moshel-api/routers"
)

func main() {
	r := router.CreateRoutes()

	r.Run(":8080")
}

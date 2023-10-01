package main

import (
	"Assignment-2/database"
	"Assignment-2/routers"
	"log"
)

func main() {
	database.StartDB()

	// Create a Gin router
	routes := routers.Router()
	log.Fatal(routes.Run(":8080"))
}

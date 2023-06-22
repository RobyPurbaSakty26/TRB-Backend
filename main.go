package main

import (
	"log"
	"os"
	"trb-backend/router"
)

func main() {

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Set a default port if "PORT" environment variable is not set
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

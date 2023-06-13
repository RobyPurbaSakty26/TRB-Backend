package main

import (
	"log"
	"trb-backend/router"
)

func main() {

	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

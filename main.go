package main

import (
	"log"

	"codedolphin.io/users-service/models"
	"codedolphin.io/users-service/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.Init()
	routes.Run()
}

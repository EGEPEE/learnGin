package main

import (
	"log"

	"github.com/EGEPEE/learnGin/delivery/restapi"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/joho/godotenv"
)

func init() {
	repository.Open()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := restapi.SetupRouter()
	// running
	r.Run()
}

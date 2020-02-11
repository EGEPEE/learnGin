package main

import (
	"log"
	"net/http"

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

	e := restapi.SetupRouter()

	if err := http.ListenAndServe(":8000", e); err != nil {
		log.Fatal(err)
	}

}

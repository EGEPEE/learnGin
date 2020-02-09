package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/EGEPEE/learnGin/delivery/restapi"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func init() {
	repository.Open()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sa := option.WithCredentialsFile("./repository/firebase/ServiceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	r := restapi.SetupRouter()
	// running
	r.Run()
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EGEPEE/learnGin/delivery/helper"
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

	iv := helper.GetIV()
	enc := helper.GCM_encrypt(os.Getenv("ENC_PWD"), "aryadanu123", iv, []byte(os.Getenv("ADD_AES")))
	dct := helper.GCM_decrypt(os.Getenv("ENC_PWD"), enc, iv, []byte(os.Getenv("ADD_AES")))
	str := string(iv[:])
	arr := []byte(str)
	fmt.Println("iv :", iv)
	fmt.Println("str :", str)
	fmt.Println("arr :", arr)
	fmt.Println("decrypt :", dct)
	r := restapi.SetupRouter()
	// running
	r.Run()
}

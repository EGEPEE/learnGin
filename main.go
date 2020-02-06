package main

import (
	"github.com/EGEPEE/learnGin/delivery/restapi"
	"github.com/EGEPEE/learnGin/repository"
)

func init() {
	repository.Open()
}

func main() {

	r := restapi.SetupRouter()
	// running
	r.Run()
}

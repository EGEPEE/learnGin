package main

import (
	"github.com/EGEPEE/learnGin/controllers"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
)

func init() {
	repository.Open()
}

func main() {
	router := gin.Default()

	userapi := router.Group("/api/usr_userapi")
	{
		userapi.GET("/get_user", controllers.GetUser)
	}

	router.Run()
}

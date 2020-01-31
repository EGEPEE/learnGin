package main

import (
	"github.com/EGEPEE/learnGin/controllers"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	repository.Open()
}

func main() {
	router := gin.Default()

	userapi := router.Group("/api/usr_userapi")
	{
		userapi.POST("/check_phonenumber", controllers.CheckPhoneNumber)
	}

	router.Run()
}

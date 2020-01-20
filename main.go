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

	v1 := router.Group("/api/v1/userapi")
	{
		v1.GET("/", controllers.FetchAllUser)
		v1.POST("/create", controllers.CreateUser)
	}

	router.Run()
}

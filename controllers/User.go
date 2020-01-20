package controllers

import (
	"net/http"
	"strconv"

	"github.com/EGEPEE/learnGin/migrations"
	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"

	"github.com/gin-gonic/gin"
)

func FetchAllUser(c *gin.Context) {
	var user []migrations.User
	var _user []models.User

	repository.DB.Find(&user)

	if len(user) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	for _, item := range user {
		_user = append(_user, models.User{ID: item.ID, Name: item.Name})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _user})
}

func CreateUser(c *gin.Context) {
	age, _ := strconv.Atoi(c.PostForm("age"))
	user := models.User{Name: c.PostForm("name"), Age: age}
	repository.DB.Save(&user)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": user.ID})
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/EGEPEE/learnGin/migrations"
	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"

	"github.com/gin-gonic/gin"
)

func FetchAll(c *gin.Context) {
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

func Create(c *gin.Context) {
	age, _ := strconv.Atoi(c.PostForm("age"))
	user := models.User{Name: c.PostForm("name"), Age: age, MemberNumber: c.PostForm("membernumber"), Email: c.PostForm("email")}

	if err := repository.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Failed", "alert": err})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "True", "alert": "User created successfully!", "resourceId": user.ID})
}

func Update(c *gin.Context) {
	var user models.User
	userID := c.PostForm("id")

	repository.DB.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User Not Found."})

		return
	}

	age, _ := strconv.Atoi(c.PostForm("age"))
	repository.DB.Model(&user).Update(models.User{Name: c.PostForm("name"), Age: age, MemberNumber: c.PostForm("membernumber")})

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

func Delete(c *gin.Context) {
	var user models.User
	userID := c.PostForm("id")

	repository.DB.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User Not Found."})

		return
	}
	repository.DB.Delete(&user)

	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "message": "User deleted."})
}

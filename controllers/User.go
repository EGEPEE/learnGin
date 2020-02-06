package controllers

import (
	"net/http"

	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []models.CustomerMain
	err := repository.GetUser(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "data": user})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": user})
}

func CheckPhone(c *gin.Context) {
	var userCP models.CustomerCheckPhone
	err := repository.CheckPhone(&userCP)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "data": userCP})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "id": userCP.CustomerMain.Name, "role_user": userCP.RoleUser, "otp_input": userCP.OtpInput})
}

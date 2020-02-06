package controllers

import (
	"net/http"

	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []models.CustomerMain
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": user})
}

func CheckPhoneNumber(c *gin.Context) {
	var user models.CustomerDetail
	noTelepon := c.PostForm("no_telepon")

	if err := repository.DB.Table("tab_master_customers").Where("no_telepon = ?", noTelepon).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})
	
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": user})
	}
	
	func CheckPhoneNumber(c *gin.Context) {
		var user models.CustomerDetail
		noTelepon := c.PostForm("no_telepon")
	
		if err := repository.DB.Table("tab_master_customers").Where("no_telepon = ?", noTelepon).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "True"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "False", "no_telepon": user.NoTelepon, "role_user": user.RoleUser, "otp_input": user.OtpInput})
}

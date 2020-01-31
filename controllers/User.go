package controllers

import (
	"fmt"
	"net/http"

	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
)

func CheckPhoneNumber(c *gin.Context) {
	var user models.TabMasterCustomer
	noTelepon := c.PostForm("no_telepon")
	repository.DB.Where("no_telepon = ?", noTelepon).First(&user)
	fmt.Printf(user.NoTelepon)

	if len(user.NoTelepon) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "True"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "False", "id": user.Name, "no_telepon": user.NoTelepon, "role_user": user.RoleUser, "otp_input": user.OtpInput})
}

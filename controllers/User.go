package controllers

import (
	"net/http"
	"os"

	"github.com/EGEPEE/learnGin/delivery/helper"
	"github.com/EGEPEE/learnGin/models"
	"github.com/EGEPEE/learnGin/repository"
	"github.com/gin-gonic/gin"
)

func GetAllAcount(c *gin.Context) {
	var user []models.CustomerMain
	err := repository.GetAllAcount(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "data": user})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": user})
}

func CheckPhone(c *gin.Context) {
	var user models.CustomerCheckPhone
	noTelepon := c.PostForm("no_telepon")
	err := repository.CheckPhone(&user, noTelepon)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "alert": "Tidak ada data."})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "id": user.CustomerMain.Name, "role_user": user.RoleUser, "otp_input": user.OtpInput})
}

func DeleteAccount(c *gin.Context) {
	// Cek apakah data ada atau tidak
	var userCheck models.CustomerCheckPhone
	noTelepon := c.PostForm("no_telepon")
	check_data := repository.CheckPhone(&userCheck, noTelepon)

	if check_data != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "alert": "Not Found"})

		return
	}

	// Delete data
	var user models.CustomerMain
	err := repository.DeleteUser(&user, noTelepon)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "alert": "Gagal menghapus data."})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "alert": "Data berhasil dihapus."})
}

func CheckPin(c *gin.Context) {
	var user models.CustomerPrivate
	noTelepon := c.PostForm("no_telepon")
	pin := c.PostForm("pin")

	// using hasing cipher keys
	pin = helper.Encrypt([]byte(pin), os.Getenv("ENV_ENCR"))

	err := repository.CheckPin(&user, noTelepon, pin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found", "data": "Pin salah"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": "Pin benar"})
}

func SetPin(c *gin.Context) {
	var user models.CustomerPrivate
	noTelepon := c.PostForm("no_telepon")
	pin := c.PostForm("pin")
	pin = helper.Encrypt([]byte(pin), os.Getenv("ENV_ENCRYPT"))

	err := repository.SetPin(&user, noTelepon, pin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success"})
}

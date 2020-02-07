package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/EGEPEE/learnGin/delivery/helper"
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

func Register(c *gin.Context) {
	cusMain := models.CustomerMain{Nama: c.PostForm("nama"), Name: "BS-00100", NoTelepon: c.PostForm("no_telepon")}
	kec := c.PostForm("kecamatan")
	kec = strings.NewReplacer("Kota ", "", "Kecamatan ", "").Replace(kec)
	pwd := c.PostForm("meta_data")
	pwd = helper.GCM_encrypt(os.Getenv("ENC_PWD"), pwd, []byte(os.Getenv("ADD_AES")))
	cusPwd := models.CustomerPassword{MetaData: pwd}
	regis := models.CustomerRegister{CustomerMain: cusMain, CustomerPassword: cusPwd, Kecamatan: kec, TanggalLahir: c.PostForm("tanggal_lahir"), Latlong: c.PostForm("latlong"), UnitDefault: c.PostForm("unit_default"), TokenFb: c.PostForm("token_fb"), NamaSupplier: c.PostForm("nama_supplier"), RoleUser: c.PostForm("role_user"), OtpInput: c.PostForm("otp_input")}
	c.BindJSON(&regis)
	err := repository.UserRegister(&regis)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True"})
}

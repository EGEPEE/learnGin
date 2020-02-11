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
	var user models.CustomerCheck
	noTelepon := c.PostForm("no_telepon")
	err := repository.CheckPhone(&user, noTelepon)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False", "role_user": user.RoleUser, "otp_input": user.OtpInput})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "role_user": user.RoleUser, "otp_input": user.OtpInput})
}

func DeleteAccount(c *gin.Context) {
	// Cek apakah data ada atau tidak
	var userCheck models.CustomerCheck
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

	err := repository.CheckPin(&user, noTelepon, pin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found", "data": "Pin salah"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "data": "Pin benar"})
}

func SetPin(c *gin.Context) {
	noTelepon := c.PostForm("no_telepon")
	pin := c.PostForm("pin")
	pin = helper.GCM_encrypt(os.Getenv("ENC_PWD"), pin, []byte(os.Getenv("ADD_AES")))

	cusPin := models.CustomerPrivate{Pin: pin}

	c.BindJSON(&cusPin)

	err := repository.SetPin(&cusPin, noTelepon)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not Found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success"})
}

func Register(c *gin.Context) {
	cusMain := models.CustomerMain{Nama: c.PostForm("nama"), Name: c.PostForm("name"), NoTelepon: c.PostForm("no_telepon")}

	kec := c.PostForm("kecamatan")
	kec = strings.NewReplacer("Kota ", "", "Kecamatan ", "").Replace(kec)

	pwd := c.PostForm("meta_data")
	pwd = helper.GCM_encrypt(os.Getenv("ENC_PWD"), pwd, []byte(os.Getenv("ADD_AES")))
	cusPwd := models.CustomerPrivate{MetaData: pwd}

	regis := models.CustomerRegister{CustomerMain: cusMain, CustomerPrivate: cusPwd, Kecamatan: kec, TanggalLahir: c.PostForm("tanggal_lahir"), Latlong: c.PostForm("latlong"), UnitDefault: c.PostForm("unit_default"), TokenFb: c.PostForm("token_fb"), NamaSupplier: c.PostForm("nama_supplier")}

	c.BindJSON(&regis)
	err := repository.UserRegister(&regis)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	setOtp := models.CustomerCheck{RoleUser: c.PostForm("role_user"), OtpInput: "2"}

	err = repository.SetOtpInput(&setOtp, c.PostForm("no_telepon"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True"})
}

func ValidateOtp(c *gin.Context) {
	var otp models.CustomerCheckOtp

	noTelepon := c.PostForm("no_telepon")
	otpGenerate := c.PostForm("otp_generate")

	err := repository.CheckOtp(&otp, noTelepon, otpGenerate)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	setOtp := models.CustomerCheck{OtpInput: "1"}

	err = repository.SetOtpInput(&setOtp, noTelepon)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "False"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "True", "no_telepon": noTelepon})
}

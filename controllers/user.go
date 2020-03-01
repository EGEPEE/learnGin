package controllers

import (
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EGEPEE/learnGin/delivery/helper"
	"github.com/EGEPEE/learnGin/models"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func (a *models.UserPrivate) CheckUser(bool, error) {
	return models.CheckUser(a.Username, util.EncodeMD5(a.Password))
}

func Auth(c *gin.Context) {

	appG := helper.Gin{C: c}
	var reqInfo models.Auth
	err := c.BindJSON(&reqInfo)

	if err != nil {
		appG.Response(http.StatusBadRequest, helper.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.NoTelepon, 13, "no_telepon").Message("Hingga 13 karakter")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("Hingga 100 karakter")

	if valid.HasErrors() {
		helper.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, helper.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	userMain := models.UserMain{NoTelepon: reqInfo.NoTelepon}
	userService := models.UserPrivate{UserMain: userMain, MetaData: reqInfo.Password}
	isExist, err := CheckUser(&userService)
}

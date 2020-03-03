package repository

import (
	"github.com/EGEPEE/learnGin/models"
	"github.com/jinzhu/gorm"
)

func CheckUser(noTelepon, password string) (bool, error) {
	var user models.UserPrivate

	userMain := models.UserMain{NoTelepon: noTelepon}
	err := db.Select("id").Where(models.UserPrivate{UserMain: userMain, MetaData: password}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if len(err) > 0 {
		return true, nil
	}

	return false, nil
}

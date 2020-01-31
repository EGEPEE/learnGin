package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Open() error {
	var err error
	DB, err = gorm.Open("mysql", "admin:password@/digitalwastev2?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return err
	}

	// DB.AutoMigrate(&migrations.User{})
	return err
}

func Close() error {
	return DB.Close()
}

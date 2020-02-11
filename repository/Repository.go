package repository

import (
	"github.com/EGEPEE/learnGin/models"
)

var nameTable = map[string]string{
	"masterCustomer": "tab_master_customers",
}

func GetAllAcount(c *[]models.CustomerMain) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, nama, name").Find(&c).Error; err != nil {
		return err
	}

	return nil
}

func CheckPhone(c *models.CustomerCheck, noTelepon string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, nama, name, otp_input, role_user").Where("no_telepon = ?", noTelepon).First(&c).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(c *models.CustomerMain, noTelepon string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Where("no_telepon = ?", noTelepon).Delete(c).Error; err != nil {
		return err
	}

	return nil
}

func CheckPrivate(c *models.CustomerPrivate, noTelepon string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, pin").Where("no_telepon = ?", noTelepon).First(&c).Error; err != nil {
		return err
	}

	return nil
}

func UserRegister(c *models.CustomerRegister) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Create(&c).Error; err != nil {
		return err
	}

	return nil
}

func SetPin(c *models.CustomerPrivate, noTelepon string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Where("no_telepon = ?", noTelepon).Update(&c).Error; err != nil {
		return err
	}

	return nil
}

func CheckOtp(c *models.CustomerCheckOtp, noTelepon, otpGenerate string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, otp_generate").Where("no_telepon = ? AND otp_generate = ?", noTelepon, otpGenerate).First(&c).Error; err != nil {
		return err
	}

	return nil
}

func SetOtpInput(c *models.CustomerCheck, noTelepon string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Where("no_telepon = ?", noTelepon).Update(&c).Error; err != nil {
		return err
	}

	return nil
}

func CustomerCheckPrivate(c *models.CustomerPrivate) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Where(&c).First(&c).Error; err != nil {
		return err
	}

	return
}

package repository

import "github.com/EGEPEE/learnGin/models"

var nameTable = map[string]string{
	"masterCustomer": "tab_master_customers",
}

func GetAllAcount(c *[]models.CustomerMain) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, nama, name").Find(&c).Error; err != nil {
		return err
	}

	return nil
}

func CheckPhone(c *models.CustomerCheckPhone, noTelepon string) (err error) {
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

func CheckPin(c *models.CustomerPrivate, noTelepon, pin string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, pin").Where("pin = ? and no_telepon = ?", pin, noTelepon).Error; err != nil {
		return err
	}

	return nil
}

func SetPin(c *models.CustomerPrivate, noTelepon, pin string) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Where("no_telepon = ? and pin = ?", noTelepon, pin).Update("pin", pin).Error; err != nil {
		return err
	}

	return nil
}

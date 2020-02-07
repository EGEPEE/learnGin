package repository

import "github.com/EGEPEE/learnGin/models"

var nameTable = map[string]string{
	"masterCustomer": "tab_master_customers",
}

func GetUser(c *[]models.CustomerMain) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, nama, name").Find(&c).Error; err != nil {
		return err
	}

	return nil
}

func CheckPhone(c *models.CustomerCheckPhone) (err error) {
	if err := DB.Table(nameTable["masterCustomer"]).Select("no_telepon, nama, name, otp_input, role_user").Find(&c).Error; err != nil {
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

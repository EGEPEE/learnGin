package repository

import "github.com/EGEPEE/learnGin/models"

func GetUser(c *[]models.CustomerMain) (err error) {
	if err := DB.Table("tab_master_customers").Select("no_telepon, alamat, name, creation, modified").Find(&c).Error; err != nil {
		return err
	}

	return nil
}

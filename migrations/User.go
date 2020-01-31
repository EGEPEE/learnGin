package migrations

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	Name     string    `gorm:"primary_key" json:"name"`
	Creation time.Time `json:"creation"`
	Modified time.Time `json:"modified"`
}

type User struct {
	gorm.Model
	ModifiedBy   string    `gorm:"type:varchar(140)" json:"modified_by"`
	Owner        string    `gorm:"type:varchar(140)" json:"owner"`
	Docstatus    int       `json:"docstatus"`
	Parent       string    `gorm:"type:varchar(140)" json:"parent"`
	Parentfield  string    `gorm:"type:varchar(140)" json:"parentfield"`
	Parenttype   string    `gorm:"type:varchar(140)" json:"parenttype"`
	Idx          int       `json:"idx"`
	Alamat       string    `gorm:"type:text" json:"alamat"`
	Kecamatan    string    `gorm:"type:varchar(140)" json:"kecamatan"`
	Foto         string    `gorm:"type:text" json:"foto"`
	Nik          string    `gorm:"type:varchar(140)" json:"nik"`
	BankAccount  string    `gorm:"type:varchar(140)" json:"bank_account"`
	RoleUser     string    `gorm:"type:varchar(140)" json:"role_user"`
	Latlong      string    `gorm:"type:varchar(140)" json:"latlong"`
	OtpGenerate  string    `gorm:"type:varchar(140)" json:"otp_generate"`
	NoTelepon    string    `gorm:"type:varchar(140)" json:"no_telepon"`
	NamaSupplier string    `gorm:"type:varchar(140)" json:"nama_supplier"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Email        string    `gorm:"type:varchar(140)" json:"email"`
	Nama         string    `gorm:"type:varchar(140)" json:"nama"`
	UnitDefault  string    `gorm:"type:varchar(140)" json:"unit_default"`
	NamingSeries string    `gorm:"type:varchar(140)" json:"naming_series"`
	MetaData     string    `gorm:"type:varchar(140)" json:"meta_data"`
	OtpInput     string    `gorm:"type:varchar(140)" json:"otp_input"`
	TokenFb      string    `gorm:"type:text" json:"token_fb"`
	Pin          string    `gorm:"type:varchar(140)" json:"pin"`
	Data29       string    `gorm:"type:varchar(140)" json:"data_29"`
	AndroID      string    `gorm:"type:varchar(140)" json:"andro_id"`
}

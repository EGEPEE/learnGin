package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CustomerMain struct {
	Name      string `json:"name"`
	Nama      string `json:"nama"`
	NoTelepon string `json:"no_telepon"`
}

type CustomerDetail struct {
	BankAccount *string `json:"bank_account"`
	CustomerRegister
}

type CustomerRegister struct {
	CustomerMain
	CustomerPrivate
	gorm.Model
	Alamat       string `json:"alamat"`
	Kecamatan    string `json:"kecamatan"`
	Latlong      string `json:"latlong"`
	TanggalLahir string `json:"tanggal_lahir"`
	UnitDefault  string `json:"unit_default"`
	TokenFb      string `json:"token_fb"`
	NamaSupplier string `json:"nama_supplier"`
}

type CustomerPrivate struct {
	CustomerMain
	Pin      string `json:"pin"`
	MetaData string `json:"meta_data"`
}

type CustomerCheck struct {
	RoleUser string `json:"role_user"`
	OtpInput string `json:"otp_input"`
}

type CustomerCheckOtp struct {
	OtpGenerate string `json:"otp_generate"`
}

package models

import (
	"time"
)

type TimeManage struct {
	Creation *time.Time `json:"creation"`
	Modified *time.Time `json:"modified"`
}

type CustomerMain struct {
	Name      *string `json:"name"`
	Nama      *string `json:"nama"`
	NoTelepon *string `json:"no_telepon"`
}

type CustomerDetail struct {
	Alamat       *string    `json:"alamat"`
	Kecamatan    *string    `json:"kecamatan"`
	Foto         *string    `json:"foto"`
	Latlong      *string    `json:"latlong"`
	TanggalLahir *time.Time `json:"tanggal_lahir"`
	Email        *string    `json:"email"`
	UnitDefault  *string    `json:"unit_default"`
	BankAccount  *string    `json:"bank_account"`
	TokenFb      *string    `json:"token_fb"`
	NamaSupplier *string    `json:"nama_supplier"`
	CustomerCheckPhone
}

type CustomerCheckPhone struct {
	RoleUser *string `json:"role_user"`
	OtpInput *string `json:"otp_input"`
	CustomerMain
}

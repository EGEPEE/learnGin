package models

import "github.com/casbin/casbin"

type Auth struct {
	Id        int    `json:"id"`
	NoTelepon string `json:"no_telepon"`
	Password  string `json:"password"`
	Role      int    `json:"role_id"`
}

type UserMain struct {
	Nama      string `json:"nama"`
	NoTelepon string `json:"no_telepon"`

	Enforcer *casbin.Enforcer `inject:""`
}

type UserRegister struct {
	UserMain
	Alamat       string `json:"alamat"`
	Kecamatan    string `json:"kecamatan"`
	Latlong      string `json:"latlong"`
	TanggalLahir string `json:"tanggal_lahir"`
	UnitDefault  string `json:"unit_default"`
	TokenFb      string `json:"token_fb"`
	NamaSupplier string `json:"nama_supplier"`
}

type UserPrivate struct {
	UserMain
	Pin      string `json:"pin"`
	MetaData string `json:"meta_data"`
}

package models

import (
	_ "github.com/EGEPEE/learnGin/repository"
)

type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	MemberNumber string `json:"membernumber"`
	Email        string `json:"email"`
}

type Repository interface {
	FetchAllUser()
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NIP 		string 		`json:"nip" gorm:"unique"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email" gorm:"unique"`
	Password 	string 		`json:"password"`
	Presensis 	[]Presensi
}

type CreateUserInput struct {
	NIP 		string `json:"nip" binding:"required"`
	Name 		string `json:"name" binding:"required"`
	Email 		string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}

type UserLoginInput struct {
	Email 		string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}
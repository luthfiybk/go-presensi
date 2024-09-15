package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NIP 		string 		`json:"nip" gorm:"unique"`
	Name 		string 		`json:"name"`
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
	Presensi 	[]Presensi
}

type CreateUserInput struct {
	Name 		string `json:"name" binding:"required"`
	Email 		string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Presensi struct {
	gorm.Model
	NIP			string		`json:"nip"`
	JamMasuk 	time.Time	`json:"jam_masuk"`
	Status 		string 		`json:"status"`
}

type CreatePresensiInput struct {
	NIP			string	`json:"nip" binding:"required"`
}
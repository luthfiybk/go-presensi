package models

import (
	"time"

	"gorm.io/gorm"
)

type Presensi struct {
	gorm.Model
	UserID		uint		`gorm:"user_id"`
	NIP			string		`json:"nip"`
	JamMasuk 	time.Time	`json:"jam_masuk"`
	Status 		string 		`json:"status"`
}

type CreatePresensiInput struct {
	UserID		uint	`json:"user_id" binding:"required"`
}
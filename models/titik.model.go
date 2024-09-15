package models

import "gorm.io/gorm"

type Titik struct {
	gorm.Model
	NamaTitik 	string 		`json:"nama_titik"`
	Latitude 	float64 	`json:"latitude"`
	Longitude 	float64 	`json:"longitude"`
	Radius 		int 		`json:"radius"`
}

type CreateTitikInput struct {
	NamaTitik 	string 		`json:"nama_titik" binding:"required"`
	Latitude 	float64 	`json:"latitude" binding:"required"`
	Longitude 	float64		`json:"longitude" binding:"required"`
	Radius 		int 		`json:"radius" binding:"required"`
}

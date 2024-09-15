package controllers

import (
	"go-presensi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TitikController struct {
	DB *gorm.DB
}

func NewTItikController (DB *gorm.DB) TitikController {
	return TitikController{DB}
}

func (tc *TitikController) FindAll(ctx *gin.Context) {
	var titik []models.Titik

	error := tc.DB.Find(&titik).Error

	if error != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"message": "Success",
		"data": titik,
	})
}

func (tc *TitikController) CreateTitik(ctx *gin.Context) {
	var payload *models.CreateTitikInput

	ctx.ShouldBindJSON(&payload)

	newTitik := models.Titik{
		NamaTitik: payload.NamaTitik,
		Latitude: payload.Latitude,
		Longitude: payload.Longitude,
		Radius: payload.Radius,
	}

	result := tc.DB.Create(&newTitik)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(201, gin.H{
		"status": 201,
		"message": "Titik berhasil ditambahkan",
		"data": newTitik,
	})
}
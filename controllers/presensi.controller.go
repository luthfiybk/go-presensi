package controllers

import (
	"go-presensi/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PresensiController struct {
	DB *gorm.DB
}

func NewPresensiController(DB *gorm.DB) PresensiController {
	return PresensiController{DB}
}

func (pc *PresensiController) PresensiMasuk(ctx *gin.Context) {
	var payload *models.CreatePresensiInput

	ctx.ShouldBindJSON(&payload)

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(401, gin.H{
			"status": 401,
			"message": "Unauthorized",
		})
	}

	payload.UserID = user.(models.User).ID

	newPresensi := models.Presensi{
		UserID: payload.UserID,
		NIP: user.(models.User).NIP,
		JamMasuk: time.Now(),
		Status: "Tepat Waktu",
	}

	result := pc.DB.Create(&newPresensi)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(201, gin.H{
		"status": 201,
		"message": "Presensi berhasil",
		"data": newPresensi,
	})
}

func (pc *PresensiController) GetAll(ctx *gin.Context) {
	var presensi []models.Presensi

	pc.DB.Find(&presensi)

	ctx.JSON(200, gin.H{
		"status": 200,
		"message": "Success",
		"data": presensi,
	})
}

func (pc *PresensiController) GetPresensiByNIP(ctx *gin.Context) {
	nip := ctx.Param("nip")

	var presensi []models.Presensi

	pc.DB.Where("nip = ?", nip).Find(&presensi)

	if len(presensi) == 0 {
		ctx.JSON(404, gin.H{
			"status": 404,
			"message": "Data not found",
		})
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"message": "Success",
		"data": presensi,
	})
}
package controllers

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go-presensi/models"
)

type UserController struct {
	DB *gorm.DB
}

func NewuserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) FindAll(ctx *gin.Context) {
	// your code here
	var users []models.User

	error := uc.DB.Find(&users).Error

	if error != nil {
		// handle error
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"message": "Success",
		"data": users,
	})
}

func (uc *UserController) FindByNIP(ctx *gin.Context) {
	var user models.User

	error := uc.DB.Where("nip = ?", ctx.Param("nip")).First(&user).Error

	if error != nil {
		// handle error
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"message": "Success",
		"data": user,
	})
}
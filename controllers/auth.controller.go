package controllers

import (
	"go-presensi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var payload *models.UserLoginInput

	if ctx.ShouldBindJSON(&payload) != nil {
		ctx.JSON(400, gin.H{
			"status": 400,
			"message": "Bad Request",
		})
		return
	}

	var user models.User

	result := ac.DB.Where("email = ?", payload.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"status": 404,
			"message": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		ctx.JSON(401, gin.H{
			"status": 401,
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"nip": user.NIP,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600, "", "localhost", false, true)
}

func (ac *AuthController) Create(ctx *gin.Context) {
	var payload *models.CreateUserInput

	if ctx.ShouldBindJSON(&payload) != nil {
		ctx.JSON(400, gin.H{
			"status": 400,
			"message": "Bad Request",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Failed to hash password",
		})
		return
	}

	newUser := models.User{
		NIP: payload.NIP,
		Name: payload.Name,
		Email: payload.Email,
		Password: string(hashedPassword),
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil {
		ctx.JSON(500, gin.H{
			"status": 500,
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status": 201,
		"message": "User created successfully",
		"data": newUser,
	})
}
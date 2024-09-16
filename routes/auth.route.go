package routes

import (
	"go-presensi/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRoute(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/login", ac.authController.Login)
	router.POST("/register", ac.authController.Create)
	router.GET("/logout", ac.authController.Logout)
}
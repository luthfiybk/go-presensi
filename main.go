package main

import (
	"go-presensi/cfg"
	"go-presensi/controllers"
	"go-presensi/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	server 					*gin.Engine

	AuthController 			controllers.AuthController
	AuthRouteController 	routes.AuthRouteController

	PresensiController 		controllers.PresensiController
	PresensiRouteController routes.PresensiRouteController
)

func init() {
	cfg.LoadEnv()
	cfg.ConnectDB()

	AuthController = controllers.NewAuthController(cfg.DB)
	AuthRouteController = routes.NewAuthRoute(AuthController)

	PresensiController = controllers.NewPresensiController(cfg.DB)
	PresensiRouteController = routes.NewPresensiRoute(PresensiController)

	server = gin.Default()
}

func main() {
	cfg.LoadEnv()

	server := gin.Default()

	router := server.Group("/api")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "Welcome to Go-Presensi API",
		})
	})

	AuthRouteController.AuthRoute(router)
	PresensiRouteController.PresensiRoute(router)

	server.Run(":8080")
}
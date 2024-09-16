package routes

import (
	"go-presensi/controllers"
	"go-presensi/middlewares"

	"github.com/gin-gonic/gin"
)

type PresensiRouteController struct {
	presensiController controllers.PresensiController
}

func NewPresensiRoute(presensiController controllers.PresensiController) PresensiRouteController {
	return PresensiRouteController{presensiController}
}

func (pc *PresensiRouteController) PresensiRoute(rg *gin.RouterGroup) {
	router := rg.Group("/presensi")

	router.GET("/", pc.presensiController.GetAll)
	router.POST("/masuk", middlewares.RequireAuth, pc.presensiController.PresensiMasuk)
	router.GET("/:id", pc.presensiController.GetPresensiByNIP)
}
package main

import (
	"go-presensi/cfg"
	"go-presensi/models"
)

func init() {
	cfg.LoadEnv()
	cfg.ConnectDB()
}

func main() {
	cfg.DB.AutoMigrate(&models.Presensi{})
	cfg.DB.AutoMigrate(&models.User{})
}
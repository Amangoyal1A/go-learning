package main

import (
	"go-user-api/config"
	"go-user-api/middlewares"
	"go-user-api/models"
	"go-user-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.CustomRecovery())
	config.ConnectDatabase()

	// Auto migrate the User table
	config.DB.AutoMigrate(&models.User{})

	routes.SetupRoutes(router)

	router.Run(":9000")
}

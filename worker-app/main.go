package main

import (
	"worker-app/config"
	"worker-app/controllers"
	"worker-app/models"
	"worker-app/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	config.ConnectDB()

	// Auto migrate model
	config.DB.AutoMigrate(&models.Job{})

	// Start Worker Pool
	worker.StartWorkerPool(500)

	// Setup Gin
	r := gin.Default()
	r.POST("/job", controllers.CreateJob)

	r.Run(":8080")
}

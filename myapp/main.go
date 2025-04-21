package main

import (
	"log"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.Default() creates a Gin engine with Logger and Recovery middleware by default.
	// But here, we will use our custom recovery middleware from our own package.
	router := gin.New()

	// Global middleware: use our custom Recovery middleware.
	router.Use(routes.RecoveryMiddleware())

	// Setup routes (endpoints)
	routes.SetupRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"trace-id-logging/handler"
	"trace-id-logging/middleware"
)

func main() {
	r := gin.Default()

	// Add trace ID middleware
	r.Use(middleware.RequestIDMiddleware())

	api := r.Group("/api")
	{
		api.GET("/users/:id", handler.NewUserHandler().GetUser)
		api.POST("/users", handler.NewUserHandler().CreateUser)
	}

	log.Println("Server starting on port 8080")
	r.Run(":8080")
}

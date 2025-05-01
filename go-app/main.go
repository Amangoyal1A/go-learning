package main

import (
	"go-app/internal/handler"
	"go-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.RequestIDMiddleware())
	// r.Use(logger.L())

	r.GET("/user/:id", handler.GetUser)

	r.Run(":8080")
}

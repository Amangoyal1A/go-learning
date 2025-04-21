package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes defines all API routes
func SetupRoutes(r *gin.Engine) {
	// Public route (example: health check)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API group (with authentication middleware)
	api := r.Group("/api")
	{
		// Use Auth middleware to protect routes
		api.Use(middlewares.AuthMiddleware())

		// Route to create a user
		api.POST("/user", controllers.CreateUser)
		// Route to fetch user by id
		api.GET("/user/:id", controllers.GetUser)
	}
}

// RecoveryMiddleware is our custom recovery middleware (wraps panic and recovers)
func RecoveryMiddleware() gin.HandlerFunc {
	return middlewares.RecoveryMiddleware()
}

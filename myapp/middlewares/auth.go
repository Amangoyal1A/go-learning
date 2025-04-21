package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for an Authorization header. If missing, it aborts the request.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			// Abort if unauthorized and return error.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized - No token provided",
			})
			return
		}

		// For demonstration, set a dummy userID in the context.
		c.Set("userID", 123)

		// Continue to the next handler.
		c.Next()
	}
}

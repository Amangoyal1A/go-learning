package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware is a custom middleware to recover from panics.
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Defer a function to recover from a panic.
		defer func() {
			if err := recover(); err != nil {
				// If panic occurs, abort the request and send a JSON error response.
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		// Continue executing handlers
		c.Next()
	}
}

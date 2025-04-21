package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Recover hua panic se
				log.Printf("❌ Recovered from panic: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong! Server recovered.",
				})
				c.Abort()
			}
		}()
		c.Next() // Continue to next handler
	}
}

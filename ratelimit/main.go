package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type ClientData struct {
	Requests    int
	WindowStart time.Time
}

var rateLimitStore = make(map[string]*ClientData)
var mu sync.Mutex

const (
	requestLimit  = 10
	timeWindow    = time.Minute // 1 minute
)

// Middleware for rate limiting
func rateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		mu.Lock()
		data, exists := rateLimitStore[clientIP]
		now := time.Now()

		if !exists || now.Sub(data.WindowStart) > timeWindow {
			// First request or window expired
			rateLimitStore[clientIP] = &ClientData{
				Requests:    1,
				WindowStart: now,
			}
		} else {
			if data.Requests >= requestLimit {
				mu.Unlock()
				c.JSON(http.StatusTooManyRequests, gin.H{
					"error": "Rate limit exceeded. Try again later ⏱️",
				})
				c.Abort()
				return
			}
			data.Requests++
		}
		mu.Unlock()

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(rateLimiterMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong 🏓",
		})
	})

	r.Run(":8080") // http://localhost:8080/ping
}

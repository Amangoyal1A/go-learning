package middleware

import (
    "time"

    "go-app/internal/logger"

    "github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        c.Next()

        duration := time.Since(start)
        log := logger.L(c)

        log.Infow("Request completed",
            "method", c.Request.Method,
            "path", c.Request.URL.Path,
            "status", c.Writer.Status(),
            "duration", duration.String(),
        )
    }
}

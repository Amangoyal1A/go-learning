package middleware

import (
	"go-app/internal/logger"
	"go-app/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestIDMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        reqID := c.GetHeader("X-Request-ID")
        if reqID == "" {
            reqID = uuid.New().String()
        }
        c.Set(util.ContextRequestIDKey, reqID)
        c.Writer.Header().Set("X-Request-ID", reqID)

        // Properly create child logger using zap.String()
        childLogger := logger.BaseLogger().With(
            zap.String("request_id", reqID),
        ).Sugar()

        c.Set(util.ContextLoggerKey, childLogger)

        c.Next()
    }
}

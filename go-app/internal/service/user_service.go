package service

import (
    "go-app/internal/logger"

    "github.com/gin-gonic/gin"
)

func GetUserName(c *gin.Context, id string) string {
    logger.L(c).Infow("Fetching user", "user_id", id)
    return "John Doe"
}

package logger

import (
	"go-app/internal/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var zl *zap.Logger

func init() {
    l, _ := zap.NewProduction()
    zl = l
}

func BaseLogger() *zap.Logger {
    return zl
}

func L(c *gin.Context) *zap.SugaredLogger {
    if c == nil {
        return zl.Sugar()
    }
    if lg, exists := c.Get(util.ContextLoggerKey); exists {
        if sugaredLogger, ok := lg.(*zap.SugaredLogger); ok {
            return sugaredLogger
        }
    }
    return zl.Sugar()
}

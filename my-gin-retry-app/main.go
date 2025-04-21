package main

import (
	"my-gin-retry-app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/external-retry", handler.ExternalRetryHandler)
	r.Run(":8080")
}

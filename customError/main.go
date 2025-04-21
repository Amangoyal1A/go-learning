package main

import (
	"yourapp/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/withdraw", handler.WithdrawHandler)

	router.Run(":8080")
}

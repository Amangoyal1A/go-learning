package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nested-app/config"
	"nested-app/db"
	"nested-app/internal/routes"
)

func main() {
	db.Init() // initialize DB
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Println("Server started at :8080")
	r.Run(":8080")
}
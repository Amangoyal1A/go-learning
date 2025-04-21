package main

import (
	"log"
	"nested-app/internal/routes"
	"nested-app/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	migrations.Init() // initialize DB
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Println("Server started at :8080")
	r.Run(":8080")
}

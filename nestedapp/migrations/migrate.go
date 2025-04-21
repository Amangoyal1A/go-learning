package migrations

import (
	"nested-app/config"
	"nested-app/internal/models"
)

func Init() {
	db := config.Connect()
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
package migrations

import (
	"nested-app/config"
	"nested-app/internal/model"
)

func Init() {
	db := config.Connect()
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}
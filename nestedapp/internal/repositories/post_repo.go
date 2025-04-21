package repositories

import (
	"nested-app/config"
	"nested-app/internal/models"
)

type PostRepository interface {
	Create(post *models.Post) error
}

type postRepo struct{}

func NewPostRepo() PostRepository {
	return &postRepo{}
}

func (r *postRepo) Create(post *models.Post) error {
	return config.DB.Create(post).Error
}

package repositories

import (
	"nested-app/config"
	"nested-app/internal/models"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetByID(id uint) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
}

type postRepo struct{}

func NewPostRepo() PostRepository {
	return &postRepo{}
}

func (r *postRepo) Create(post *models.Post) error {
	return config.DB.Create(post).Error
}

func (r *postRepo) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	err := config.DB.Preload("Comments").First(&post, id).Error
	return &post, err
}

func (r *postRepo) Update(post *models.Post) error {
	return config.DB.Save(post).Error
}

func (r *postRepo) Delete(id uint) error {
	return config.DB.Delete(&models.Post{}, id).Error
}
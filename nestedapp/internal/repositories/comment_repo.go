package repositories

import (
	"nested-app/config"
	"nested-app/internal/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(id uint) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id uint) error
}

type commentRepo struct{}

func NewCommentRepo() CommentRepository {
	return &commentRepo{}
}

func (r *commentRepo) Create(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}

func (r *commentRepo) GetByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := config.DB.First(&comment, id).Error
	return &comment, err
}

func (r *commentRepo) Update(comment *models.Comment) error {
	return config.DB.Save(comment).Error
}

func (r *commentRepo) Delete(id uint) error {
	return config.DB.Delete(&models.Comment{}, id).Error
}

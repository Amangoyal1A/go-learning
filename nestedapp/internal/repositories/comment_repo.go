package repositories

import (
	"nested-app/config"
	"nested-app/internal/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
}

type commentRepo struct{}

func NewCommentRepo() CommentRepository {
	return &commentRepo{}
}

func (r *commentRepo) Create(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}
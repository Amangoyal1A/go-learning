package services

import (
	"nested-app/internal/models"
	"nested-app/internal/repositories"
)

type CommentService interface {
	Create(comment *models.Comment) error
}

type commentService struct {
	repo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) Create(comment *models.Comment) error {
	return s.repo.Create(comment)
}
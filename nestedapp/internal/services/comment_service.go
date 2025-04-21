package services

import (
	"nested-app/internal/model"
	"nested-app/internal/repository"
)

type CommentService interface {
	Create(comment *model.Comment) error
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) Create(comment *model.Comment) error {
	return s.repo.Create(comment)
}
package services

import (
	"nested-app/internal/models"
	"nested-app/internal/repositories"
)

type PostService interface {
	Create(post *models.Post) error
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) Create(post *models.Post) error {
	return s.repo.Create(post)
}
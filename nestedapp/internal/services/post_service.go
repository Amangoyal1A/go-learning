package services

import (
	"nested-app/internal/model"
	"nested-app/internal/repository"
)

type PostService interface {
	Create(post *model.Post) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) Create(post *model.Post) error {
	return s.repo.Create(post)
}
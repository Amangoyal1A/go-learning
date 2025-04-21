package repositories

import "nested-app/internal/model"

type PostRepository interface {
	Create(post *model.Post) error
}

type postRepo struct{}

func NewPostRepo() PostRepository {
	return &postRepo{}
}

func (r *postRepo) Create(post *model.Post) error {
	return config.DB.Create(post).Error
}

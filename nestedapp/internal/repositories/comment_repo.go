package repositories


import "nested-app/internal/model"

type CommentRepository interface {
	Create(comment *model.Comment) error
}

type commentRepo struct{}

func NewCommentRepo() CommentRepository {
	return &commentRepo{}
}

func (r *commentRepo) Create(comment *model.Comment) error {
	return config.DB.Create(comment).Error
}
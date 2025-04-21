package services

import (
	"nested-app/internal/model"
	"nested-app/internal/repository"
)

type UserService interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

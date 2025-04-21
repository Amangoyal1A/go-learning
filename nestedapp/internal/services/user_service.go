package services

import (
	"nested-app/internal/models"
	"nested-app/internal/repositories"
)

type UserService interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

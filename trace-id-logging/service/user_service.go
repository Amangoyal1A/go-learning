// service/user_service.go
package service

import (
	"github.com/google/uuid"
	"trace-id-logging/models"
	"trace-id-logging/repository"
)

type UserService interface {
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetUser(id string) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userServiceImpl) CreateUser(user models.User) (models.User, error) {
	user.ID = uuid.New().String()
	err := s.repo.Save(user)
	return user, err
}

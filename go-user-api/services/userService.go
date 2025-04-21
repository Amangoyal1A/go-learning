package services

import (
	"go-user-api/config"
	"go-user-api/dto"
	"go-user-api/models"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error)
	GetAllUsers() ([]dto.UserResponse, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return dto.UserResponse{}, result.Error
	}
	return dto.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *userService) GetAllUsers() ([]dto.UserResponse, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, dto.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email})
	}
	return response, nil
}

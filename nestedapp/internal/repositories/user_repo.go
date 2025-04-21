package repositories

import (
	"nested-app/config"
	"nested-app/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepo struct{}

func NewUserRepo() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepo) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.Preload("Posts.Comments").First(&user, id).Error
	return &user, err
}

func (r *userRepo) Update(user *models.User) error {
	return config.DB.Save(user).Error
}

func (r *userRepo) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
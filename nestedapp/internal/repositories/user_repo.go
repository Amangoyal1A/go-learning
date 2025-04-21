package repositories

import (
	"gorm.io/gorm"
	"myapp/models"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetUser(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Posts.Comments").First(&user, id).Error
	return &user, err
}

func (r *userRepo) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
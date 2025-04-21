// repository/user_repo.go
package repository

import (
	"trace-id-logging/models"
	"errors"
)

var mockDB = map[string]models.User{}

type UserRepository interface {
	GetByID(id string) (models.User, error)
	Save(user models.User) error
}

type userRepoImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepoImpl{}
}

func (r *userRepoImpl) GetByID(id string) (models.User, error) {
	user, ok := mockDB[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepoImpl) Save(user models.User) error {
	mockDB[user.ID] = user
	return nil
}

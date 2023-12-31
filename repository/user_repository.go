package repository

import (
	"context"

	"github.com/ronnachate/inventory-api-go/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetUsers(c context.Context, page int, rows int) ([]domain.User, error) {
	var users []domain.User
	result := ur.DB.Model(&domain.User{}).Preload("Status").Scopes(NewPagination(page, rows).PageResult).Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User

	result := ur.DB.Model(&domain.User{}).Preload("Status").First(&user, "id = ?", id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

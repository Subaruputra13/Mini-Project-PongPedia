package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginUser(user *models.User) error
	CheckEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (a *authRepository) CheckEmail(email string) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *authRepository) LoginUser(user *models.User) error {

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}

	return nil
}

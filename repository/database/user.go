package database

import (
	"PongPedia/config"
	"PongPedia/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetuserByEmail(email string) (*models.User, error)
	LoginUser(email, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetuserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) LoginUser(email, password string) error {
	var user models.User

	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return err
	}

	if user.Password != password {
		return errors.New("Wrong password")
	}

	return nil
}

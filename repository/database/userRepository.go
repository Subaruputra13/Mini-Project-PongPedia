package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAlltUser() (user []models.User, err error)
	GetuserByEmail(email string) (*models.User, error)
	GetUseByIdWithCookie(id int) (*models.User, error)
	UpdateUserWithCookie(id int, user *models.User) error
	CreateUser(user *models.User) error
	ReadToken(id int) (*models.User, error)
	DeleteUser(id int, password string) (user *models.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) GetAlltUser() (user []models.User, err error) {
	if err := config.DB.Preload("Player").Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUseByIdWithCookie(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).Preload("Player").First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) ReadToken(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) UpdateUserWithCookie(id int, user *models.User) error {

	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
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

func (u *userRepository) CreateUser(user *models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteUser(id int, password string) (user *models.User, err error) {

	if err := config.DB.Where("id = ? AND password = ?", id, password).First(&user).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

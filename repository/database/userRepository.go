package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ReadToken(id int) (user *models.User, err error)
	GetUser() (user []models.User, err error)
	GetuserByEmail(email string) (*models.User, error)
	GetUseById(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}

}

func (u *userRepository) ReadToken(id int) (user *models.User, err error) {

	err = config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUser() (user []models.User, err error) {
	if err := config.DB.Preload("Player.Participation").Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUseById(id int) (user *models.User, err error) {
	err = config.DB.Model(&user).Preload("Player.Participation").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user *models.User) error {

	if err := config.DB.Updates(&user).Error; err != nil {
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

func (u *userRepository) DeleteUser(user *models.User) error {

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

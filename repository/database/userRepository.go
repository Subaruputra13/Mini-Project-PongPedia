package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() (user []models.User, err error)
	GetuserByEmail(email, username string) (user *models.User, err error)
	GetUserById(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(user *models.User) error
	CountUser() (res int64)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}

}

func (u *userRepository) GetUser() (user []models.User, err error) {
	if err := config.DB.Preload("Player.Participation").Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUserById(id uint) (user *models.User, err error) {

	if err = config.DB.Model(&user).Preload("Player.Participation.Turnament").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) CountUser() (res int64) {
	res = 0
	user := []models.User{}

	if err := config.DB.Model(&user).Count(&res).Error; err != nil {
		return 0
	}
	return res
}

func (u *userRepository) UpdateUser(user *models.User) error {

	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetuserByEmail(email, username string) (user *models.User, err error) {

	if err := config.DB.Where("email = ? OR username = ?", email, username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
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

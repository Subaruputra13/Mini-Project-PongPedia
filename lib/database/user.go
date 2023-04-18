package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

func GetUser() ([]models.User, error) {
	users := []models.User{}

	err := config.DB.Preload("Player").Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	return users, err

}

func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func LoginUser(users models.User) (models.User, error) {

	if err := config.DB.Where("nama = ? AND password = ? AND role = ?", users.Nama, users.Password, "PLAYER").Find(&users).Error; err != nil {
		return models.User{}, err
	}

	if err := config.DB.Where("nama = ? AND password = ? AND role = ?", users.Nama, users.Password, "ADMIN").Find(&users).Error; err != nil {
		return models.User{}, err
	}

	return users, nil
}

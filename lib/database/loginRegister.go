package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

func Login(users models.User) (models.User, error) {
	if err := config.DB.Where("username = ? AND password = ?", users.Username, users.Password).First(&users).Error; err != nil {
		if err = config.DB.Where("email = ? AND password = ?", users.Email, users.Password).First(&users).Error; err != nil {
			return models.User{}, err
		}
	}

	return users, nil
}

func Register(user models.User) (models.User, error) {

	// medefinisikan query untuk membuat data user(INSERT INTO users)
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err

	}

	return user, nil

}

package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

// Get All User
func GetUser() ([]models.User, error) {
	// membuat variable users dengan tipe data Slice dari struct User
	users := []models.User{}

	// mendefinisikan qeuery untuk mengambil semua data user
	err := config.DB.Preload("Player").Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	return users, err

}

// Get All User By Role
func GetUserByRole(role any) ([]models.User, error) {
	users := []models.User{}

	err := config.DB.Where("role = ?", role).Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	return users, err
}

// Update User By Id
func UpdateUserById(users models.User, id any) (models.User, error) {

	// mendefinisikan query untuk mengupdate data user berdasarkan id (UPDATE users SET ... WHERE id = ?)
	err := config.DB.Where("id = ?", id).Updates(&users).Error

	if err != nil {
		return models.User{}, err
	}

	return users, nil

}

// Delete User By Id
func DeleteUserById(id any) (interface{}, error) {

	// mendefinisikan query untuk menghapus data user berdasarkan id (DELETE FROM users WHERE id = ?)
	err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error

	if err != nil {
		return nil, err
	}

	return "Success Delete User by id", nil
}

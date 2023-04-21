package database

import (
	"PongPedia/config"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get All User
func GetUser() ([]models.User, error) {
	// membuat variable users dengan tipe data Slice dari struct User
	users := []models.User{}

	// mendefinisikan qeuery untuk mengambil semua data user
	err := config.DB.Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	return users, err

}

// Get All User By Role
func GetUserById(id any) (models.User, error) {
	// membuat variable users dengan tipe data Slice dari struct User
	users := models.User{}

	// mendefinisikan qeuery untuk mengambil semua data user dengan role tertentu
	err := config.DB.Where("id = ?", id).First(&users).Error

	if err != nil {
		return models.User{}, err
	}

	return users, err
}

// Create User
func CreateUser(user models.User) (models.User, error) {
	// //Validate Unique Data
	result := config.DB.Where("username = ? AND email = ? AND password = ?", user.Username, user.Email, user.Password).First(&user)
	if result.RowsAffected > 0 {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, "User ini sudah tersedia !")
	}
	result = config.DB.Where("username = ?", user.Username).Find(&user)
	if result.RowsAffected > 0 {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, " Username Sudah tersedia")
	}

	result = config.DB.Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected > 0 {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, " Email Sudah tersedia")
	}

	result = config.DB.Where("password = ?", user.Password).Find(&user)
	if result.RowsAffected > 0 {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, " Password Sudah tersedia")
	}

	// medefinisikan query untuk membuat data user(INSERT INTO users)
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err

	}

	return user, nil

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

// Login User
func LoginUser(users models.User) (models.User, error) {
	// mendefinisikan query untuk mengambil data user berdasarkan nama dan password (SELECT * FROM users WHERE nama = ? AND password = ?)
	err := config.DB.Where("email = ? AND password = ?", users.Email, users.Password).First(&users).Error

	if err != nil {
		return models.User{}, err
	}

	return users, nil
}

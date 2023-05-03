package database

import (
	"PongPedia/config"
	m "PongPedia/middleware"
	"PongPedia/models"

	"github.com/labstack/echo"
)

// Get User
func GetUser(c echo.Context) (models.User, error) {
	// membuat variable users dengan tipe data Slice dari struct User
	var users models.User

	id := m.Auth(c)

	if err := config.DB.Where("id = ?", id).Preload("Player").First(&users).Error; err != nil {
		return models.User{}, err
	}

	return users, nil
}

// Update User
func UpdateUser(c echo.Context) (models.User, error) {

	var users models.User

	username := c.Param("username")

	id := m.Auth(c)

	if err := config.DB.Where("id = ? AND username = ?", id, username).First(&users).Error; err != nil {
		return models.User{}, err
	}

	c.Bind(&users)

	if err := c.Validate(&users); err != nil {
		return models.User{}, err
	}

	if err := config.DB.Where("id = ? AND username = ?", id, username).Updates(&users).Error; err != nil {
		return models.User{}, err
	}

	return users, nil

}

func DeleteUser(c echo.Context) (models.User, error) {

	var users models.User

	username := c.Param("username")

	id := m.Auth(c)

	password := c.FormValue("Password")

	if err := config.DB.Where("id = ? AND username = ?", id, username).First(&users).Error; err != nil {
		return models.User{}, echo.NewHTTPError(400, "Unauthorize")
	}

	if err := config.DB.Where("id = ? AND password = ?", id, password).First(&users).Error; err != nil {
		return models.User{}, echo.NewHTTPError(400, "Wrong password !")
	}

	if err := config.DB.Where("id = ?", id).Delete(&users).Error; err != nil {
		return models.User{}, err
	}

	return users, nil
}

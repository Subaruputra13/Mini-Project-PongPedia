package database

import (
	"PongPedia/config"
	m "PongPedia/middleware"
	"PongPedia/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func Login(c echo.Context) (interface{}, error) {
	var users models.User

	userEmail := c.FormValue("Username/Email")

	c.Bind(&users)

	// Validate Required
	if err := c.Validate(&models.User{
		Username: userEmail,
		Email:    userEmail,
		Password: users.Password,
	}); err != nil {
		return nil, err
	}

	// Condition Login with username or email
	if err := config.DB.Where("username = ? AND password = ?", userEmail, users.Password).First(&users).Error; err != nil {
		if err = config.DB.Where("email = ? AND password = ?", userEmail, users.Password).First(&users).Error; err != nil {
			return models.User{}, err
		}
	}

	// Crete Token
	token, err := m.CreateToken(int(users.ID), users.Username, users.Role)

	// Create Cookie
	m.CreateCookie(c, token)

	if err != nil {
		return nil, err
	}

	userResponse := models.UserReponse{ID: users.ID, Username: users.Username, Email: users.Email, Token: token}

	return userResponse, nil
}

func Register(c echo.Context) (models.User, error) {
	var users models.User

	c.Bind(&users)

	// Validate Required
	if err := c.Validate(&users); err != nil {
		return models.User{}, err
	}

	// medefinisikan query untuk membuat data user(INSERT INTO users)
	err := config.DB.Create(&users).Error

	if err != nil {
		return models.User{}, err

	}

	return users, nil

}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)

	c.SetCookie(cookie)
	return nil
}

package controllers

import (
	"PongPedia/lib/database"
	m "PongPedia/middleware"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	users := models.User{}

	c.Bind(&users)

	// Validate Required
	if users.Username == "" && users.Password == "" {
		if users.Email == "" && users.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Field is Required",
			})
		}
	}

	if users.Username == "" && users.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Username or Email is Required",
		})
	}

	if users.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Password is Required",
		})
	}

	users, err := database.Login(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Fail Login",
			"Error":   err.Error(),
		})
	}

	// Crete Token
	token, err := m.CreateToken(int(users.ID), users.Username, users.Role)

	// Create Cookie
	m.CreateCookie(c, token)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserReponse{ID: users.ID, Username: users.Username, Email: users.Email, Token: token}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Login Succes!",
		Data:    userResponse,
	})
}

// Controllers untuk membuat data user
func RegisterControllers(c echo.Context) error {

	users := models.User{}

	c.Bind(&users)

	// Validate Required
	if users.Username == "" || users.Email == "" || users.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Username, Email and Password is Required",
		})
	}

	// memanggil fungsi CreateUser() yang ada di package database
	users, err := database.Register(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes Create data",
		Data:    users,
	})
}

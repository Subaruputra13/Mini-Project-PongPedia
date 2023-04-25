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
	if users.Username == "" {
		c.Validate(&users)
	} else {
		return c.JSON(http.StatusBadRequest, "Error Can't Handler")
	}

	// Memanggil Fungsi LoginUser() yang ada di package database
	users, err := database.LoginUser(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Fail Login",
			"Error":   err.Error(),
		})
	}

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

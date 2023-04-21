package controllers

import (
	"PongPedia/lib/database"
	m "PongPedia/middleware"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	users := models.User{}

	c.Bind(&users)

	// Validate Require
	if users.Email == "" && users.Password == "" {
		return c.JSON(http.StatusBadRequest, "Field Tidak Boleh Kosong !")
	} else if users.Email == "" || users.Password == "" {
		return c.JSON(http.StatusBadRequest, "Field Username atau Password Tidak Boleh Kosong !")
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

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserReponse{ID: users.ID, Username: users.Username, Email: users.Email, Token: token}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Login Succes!",
		Data:    userResponse,
	})
}

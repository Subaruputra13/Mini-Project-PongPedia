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

	users, err := database.LoginUser(users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	token, err := m.CreateToken(int(users.ID), users.Nama, users.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserReponse{ID: users.ID, Nama: users.Nama, Email: users.Email, Token: token}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Login Succes!",
		Data:    userResponse,
	})

}

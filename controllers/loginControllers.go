package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {

	users, err := database.Login(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Fail Login",
			"Error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Login Succes!",
		Data:    users,
	})
}

// Controllers untuk membuat data user
func RegisterControllers(c echo.Context) error {
	_, err := database.Register(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success Register")
}

func LogoutControllers(c echo.Context) error {
	err := database.Logout(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusMovedPermanently, "Logout Succes!")
}

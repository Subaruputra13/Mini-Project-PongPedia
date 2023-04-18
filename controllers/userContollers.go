package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes get user!",
		Data:    users,
	})
}
func GetUserByIdControllers(c echo.Context) {

}

func CreateUserControllers(c echo.Context) error {
	users := models.User{}

	c.Bind(&users)

	users, err := database.CreateUser(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes Create data",
		Data:    users,
	})
}

func UpdateUserByIdControllers(c echo.Context) {

}

func DeteleUserByIdControllers(c echo.Context) {

}

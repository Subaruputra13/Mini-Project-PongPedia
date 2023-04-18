package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPlayerControllers(c echo.Context) error {
	players, err := database.GetPlayer()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes get all user!",
		Data:    players,
	})
}
func GetPlayerByIdControllers(c echo.Context) {

}

func CreatePlayerControllers(c echo.Context) error {
	players := models.Player{}

	c.Bind(&players)

	players, err := database.CreatePlayer(players)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes Create Data Players",
		Data:    players,
	})
}

func UpdatePlayerByIdControllers(c echo.Context) {

}

func DetelePlayerByIdControllers(c echo.Context) {

}

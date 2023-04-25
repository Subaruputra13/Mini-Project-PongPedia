package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetPlayersControllers(c echo.Context) error {
	players, err := database.GetPlayer()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Responses{
		Message: "Success get all players",
		Data:    players,
	})
}

func CreatePlayersControllers(c echo.Context) error {
	players := models.Player{}

	c.Bind(&players)

	players, err := database.CreatePlayer(players)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes Create Player",
		Data:    players,
	})
}

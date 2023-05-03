package controllers

import (
	"PongPedia/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func MyTurnamentControllers(c echo.Context) error {
	players, err := database.MyTurnament(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, players)
}

func CreatePlayersControllers(c echo.Context) error {
	_, err := database.CreateAndUpdatePlayer(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success Update Player")
}

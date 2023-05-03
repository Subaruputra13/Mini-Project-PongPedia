package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTurnamentControllers(c echo.Context) error {
	turnaments, err := database.GetTurnament(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "All Turnament",
		Data:    turnaments,
	})
}

func GetTurnamentDetailControllers(c echo.Context) error {
	turnaments, err := database.GetTurnamentDetail(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: fmt.Sprintf("Turnament %s", turnaments.Name),
		Data:    turnaments,
	})
}

func CreateTurnamentControllers(c echo.Context) error {
	turnaments, err := database.CreateTurnament(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Success Create Turnament",
		Data:    turnaments,
	})
}

func RegisterTurnamentControllers(c echo.Context) error {
	err := database.RegisterTurnament(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success Register Turnament")
}

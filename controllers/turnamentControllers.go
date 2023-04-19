package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Contoller Get All Turnament
func GetTurnamentControllers(c echo.Context) error {

	turnaments, err := database.GetTurnament()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes Get All Turnament",
		Data:    turnaments,
	})
}

// Create Turnament Controllers
func CreateTurnamentControllers(c echo.Context) error {
	turnaments := models.Turnament{}

	c.Bind(&turnaments)

	turnaments, err := database.Createturnament(turnaments)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Success Create Turnament",
		Data:    turnaments,
	})
}

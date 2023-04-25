package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllTurnamentControllers(c echo.Context) error {
	// mengambil data dari database
	turnaments, err := database.GetTurnamnet()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Success Get All Turnament",
		Data:    turnaments,
	})
}

func CreateTurnamentControllers(c echo.Context) error {
	// membuat variable user dengan tipe data struct Turnament
	turnament := models.Turnament{}

	// mengambil data dari request body dan menampungnya ke variable user
	c.Bind(&turnament)

	// memasukan data user ke database melalui function CreateUser
	turnament, err := database.CreateTurnament(turnament)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Success Create Turnament",
		Data:    turnament,
	})
}

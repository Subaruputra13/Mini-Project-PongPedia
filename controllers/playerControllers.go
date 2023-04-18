package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controllers untuk mengambil data player
func GetPlayerControllers(c echo.Context) error {

	// memanggil fungsi GetPlayer() yang ada di package database
	players, err := database.GetPlayer()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes get all players!",
		Data:    players,
	})
}

func CreatePlayerControllers(c echo.Context) error {
	// membuat variable players dengan tipe data struct Player dari package models
	players := models.Player{}

	// mengambil data dari client dan memasukkannya ke variable players
	c.Bind(&players)

	// memanggil fungsi CreatePlayer() yang ada di package database
	players, err := database.CreatePlayer(players)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "Succes Create Data Players",
		Data:    players,
	})
}

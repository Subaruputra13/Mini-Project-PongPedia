package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Controllers untuk mengambil data user
func GetUserControllers(c echo.Context) error {

	// memanggil fungsi GetUser() yang ada di package database
	users, err := database.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: fmt.Sprintf("Selamat Datang %v", users.Username),
		Data:    users,
	})
}

// Controllers untuk mengupdate data user berdasarkan id
func UpdateUserControllers(c echo.Context) error {
	// memanggil fungsi UpdateUserById() yang ada di package database
	users, err := database.UpdateUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: fmt.Sprintf("Succes Update User %v", users.Username),
		Data:    users,
	})
}

// Controllers untuk menghapus data user berdasarkan id
func DeteleUserControllers(c echo.Context) error {
	// Memanggil fungsi DeleteUser() yang ada dipackage database
	_, err := database.DeleteUser(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success Delete User")
}

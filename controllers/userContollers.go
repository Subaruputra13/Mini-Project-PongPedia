package controllers

import (
	"PongPedia/lib/database"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controllers untuk mengambil data user
func GetUserControllers(c echo.Context) error {

	// memanggil fungsi GetUser() yang ada di package database
	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes get user!",
		Data:    users,
	})
}

// Controllers untuk mengambil data user berdasarkan role
func GetUserByIdControllers(c echo.Context) error {

	// membuat variable userRole dengan parameter role yang dikirim dari client
	userId := c.Param("id")

	// memanggil fungsi GetUserByRole() yang ada di package database
	users, err := database.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes get user by Id",
		Data:    users,
	})
}

// Controllers untuk mengupdate data user berdasarkan id
func UpdateUserByIdControllers(c echo.Context) error {
	// membuat variable userId dengan parameter id yang dikirim dari client
	userId := c.Param("id")

	// membuat variable users dengan tipe data struct User dari package models
	users := models.User{}

	// mengambil data dari client dan memasukkannya ke variable users
	c.Bind(&users)

	// memanggil fungsi UpdateUserById() yang ada di package database
	users, err := database.UpdateUserById(users, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes Update user by Id",
		Data:    users,
	})
}

// Controllers untuk menghapus data user berdasarkan id
func DeteleUserByIdControllers(c echo.Context) error {

	// membuat variable userId dengan parameter id yang dikirim dari client
	userId := c.Param("id")

	// memanggil fungsi DeleteUserById() yang ada di package database
	_, err := database.DeleteUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Success Delete user by id",
	})
}

// Controllers untuk membuat data user
func CreateUserControllers(c echo.Context) error {
	// membuat variable users dengan tipe data struct User dari package models
	users := models.User{}

	// mengambil data dari client dan memasukkannya ke variable users
	c.Bind(&users)

	// Validate Require
	if err := c.Validate(users); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Message": "Field tidak boleh kosong !",
		})
	}

	// memanggil fungsi CreateUser() yang ada di package database
	users, err := database.CreateUser(users)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Responses{
		Message: "Succes Create data",
		Data:    users,
	})
}

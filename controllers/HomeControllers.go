package controllers

import (
	"PongPedia/config"
	"PongPedia/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeController(c echo.Context) error {
	user := models.User{}

	err := config.DB.Where("role = ?", "ADMIN").Find(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"eror": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Selama Datang Di Home",
	})

}

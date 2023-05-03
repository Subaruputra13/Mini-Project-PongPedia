package database

import (
	"PongPedia/config"
	m "PongPedia/middleware"
	"PongPedia/models"

	"github.com/labstack/echo"
)

func MyTurnament(c echo.Context) (models.Player, error) {
	var users models.User
	var players models.Player

	id := m.Auth(c)

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return models.Player{}, err
	}

	if err := config.DB.Preload("Match").Where("user_id = ?", users.ID).First(&players).Error; err != nil {
		return models.Player{}, err
	}

	return players, nil
}

func CreateAndUpdatePlayer(c echo.Context) (interface{}, error) {
	var users models.User
	var players models.Player

	id := m.Auth(c)

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Where("user_id = ?", users.ID).First(&players).Error; err != nil {
		c.Bind(&players)

		if players.Name == "" || players.Age == 0 || players.BirthDate == "" || players.Gender == "" {
			return nil, echo.NewHTTPError(400, "Field can't be empty")
		}

		// if err := c.Validate(&players); err != nil {
		// 	return nil, err
		// }

		if err := config.DB.Model(&players).Where("user_id = ?", users.ID).Save(&models.Player{
			Name:      players.Name,
			Age:       players.Age,
			BirthDate: players.BirthDate,
			Gender:    players.Gender,
			UserID:    int(users.ID),
		}).Error; err != nil {
			return nil, err
		}
	}

	return players, nil

}

package database

import (
	"PongPedia/config"
	m "PongPedia/middleware"
	"PongPedia/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetTurnament(c echo.Context) ([]models.Turnament, error) {
	var turnaments []models.Turnament

	if err := config.DB.Preload("Match").Preload("Participation").Find(&turnaments).Error; err != nil {
		return []models.Turnament{}, err
	}

	return turnaments, nil
}

func GetTurnamentDetail(c echo.Context) (models.Turnament, error) {
	var turnaments models.Turnament

	turnamentId, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Preload("Match").Preload("Participation").First(&turnaments, turnamentId).Error; err != nil {
		return models.Turnament{}, err
	}

	return turnaments, nil
}

func CreateTurnament(c echo.Context) (models.Turnament, error) {
	var turnaments models.Turnament

	c.Bind(&turnaments)

	if err := config.DB.Create(&turnaments).Error; err != nil {
		return models.Turnament{}, err
	}

	return turnaments, nil
}

func RegisterTurnament(c echo.Context) error {
	var users models.User
	var turnaments models.Turnament
	var participation models.Participation
	var players models.Player

	id := m.Auth(c)
	turnamentId, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return err
	}

	if err := config.DB.First(&turnaments, turnamentId).Error; err != nil {
		return err
	}

	if err := config.DB.Where("user_id = ?", id).First(&players).Error; err != nil {
		return echo.NewHTTPError(400, "Fill your player data first")
	}

	if err := config.DB.Where("player_id = ? AND turnament_id = ?", players.ID, turnamentId).First(&participation).Error; err == nil {
		return echo.NewHTTPError(400, "User already registered")
	}

	if err := config.DB.Model(&participation).Create(&models.Participation{

		PlayerID:    int(players.ID),
		TurnamentID: turnamentId,
	}).Error; err != nil {
		return err
	}

	return nil
}

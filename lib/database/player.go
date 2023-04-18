package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

func GetPlayer() ([]models.Player, error) {
	players := []models.Player{}

	err := config.DB.Find(&players).Error

	if err != nil {
		return []models.Player{}, err
	}

	return players, err

}

func CreatePlayer(players models.Player) (models.Player, error) {
	err := config.DB.Create(&players).Error

	if err != nil {
		return models.Player{}, err
	}

	return players, nil

}

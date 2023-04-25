package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

func GetPlayer() ([]models.Player, error) {
	players := []models.Player{}

	err := config.DB.Preload("User").Preload("Turnament").Find(&players).Error

	if err != nil {
		return []models.Player{}, err
	}

	return players, err
}

func CreatePlayer(players models.Player) (models.Player, error) {
	err := config.DB.Create(&players).Error

	for _, turnamentId := range players.TurnamentId {
		match := new(models.Match)
		match.PlayerId = int(players.ID)
		match.TurnamentId = turnamentId
		config.DB.Create(&match)
	}

	if err != nil {
		return models.Player{}, err
	}

	return players, nil
}

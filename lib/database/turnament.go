package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

// Get All Turnament
func GetTurnament() ([]models.Turnament, error) {
	turnaments := []models.Turnament{}

	err := config.DB.Preload("Match").Find(&turnaments).Error

	if err != nil {
		return []models.Turnament{}, err
	}

	return turnaments, err
}

func Createturnament(turnaments models.Turnament) (models.Turnament, error) {
	err := config.DB.Create(&turnaments).Error

	if err != nil {
		return models.Turnament{}, err
	}

	return turnaments, nil
}

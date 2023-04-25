package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

func GetTurnamnet() ([]models.Turnament, error) {
	// membuat variable users dengan tipe data Slice dari struct User
	turnaments := []models.Turnament{}

	// mendefinisikan qeuery untuk mengambil semua data user
	err := config.DB.Preload("Match").Find(&turnaments).Error

	if err != nil {
		return []models.Turnament{}, err
	}

	return turnaments, err
}

func CreateTurnament(turnament models.Turnament) (models.Turnament, error) {

	// medefinisikan query untuk membuat data user(INSERT INTO users)
	err := config.DB.Create(&turnament).Error

	if err != nil {
		return models.Turnament{}, err

	}

	return turnament, nil

}

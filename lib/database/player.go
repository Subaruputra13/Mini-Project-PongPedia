package database

import (
	"PongPedia/config"
	"PongPedia/models"
)

// Get All Player
func GetPlayer() ([]models.Player, error) {

	// membuat variable players dengan tipe data Slice dari struct Player
	players := []models.Player{}

	// mendefinisikan qeuery untuk mengambil semua data player (SELECT * FROM players)
	err := config.DB.Find(&players).Error

	if err != nil {
		return []models.Player{}, err
	}

	return players, err

}

// Create Player
func CreatePlayer(players models.Player) (models.Player, error) {

	// medefinisikan query untuk membuat data player (INSERT INTO players)
	err := config.DB.Create(&players).Error

	if err != nil {
		return models.Player{}, err
	}

	return players, nil

}

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
	err := config.DB.Preload("Match").Find(&players).Error

	if err != nil {
		return []models.Player{}, err
	}

	return players, err

}

// Get Player By ID
func GetPlayerById(id any) (models.Player, error) {
	players := models.Player{}

	// mendefinisikan query untuk mengambil data player berdasarkan id (SELECT * FROM players WHERE id = ?)
	err := config.DB.Preload("Match").Where("id = ?", id).First(&players).Error

	if err != nil {
		return models.Player{}, err
	}

	return players, nil
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

// Update Player
func UpdatePlayerById(players models.Player, id any) (models.Player, error) {
	err := config.DB.Where("id = ?", id).Updates(&players).Error

	if err != nil {
		return models.Player{}, err
	}

	return players, nil
}

// Delete Player
func DeletePlayerId(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Player{}).Error

	if err != nil {
		return nil, err
	}

	return "Success Delete Player by Id", nil
}

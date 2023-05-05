package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type PlayerRespository interface {
	CreatePlayer(player *models.Player) error
	UpdatePlayer(player *models.Player) error
	GetPlayerId(id int) (*models.Player, error)
}

type playerRespository struct {
	db *gorm.DB
}

func NewPlayerRespository(db *gorm.DB) *playerRespository {
	return &playerRespository{db}
}

func (p *playerRespository) GetPlayerId(id int) (*models.Player, error) {
	var player models.Player

	if err := config.DB.Where("user_id = ?", id).Preload("Participation").First(&player).Error; err != nil {
		return nil, err
	}

	return &player, nil
}

func (p *playerRespository) CreatePlayer(player *models.Player) error {

	if err := config.DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}

func (p *playerRespository) UpdatePlayer(player *models.Player) error {

	if err := config.DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}

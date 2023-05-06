package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type PlayerRespository interface {
<<<<<<< HEAD
<<<<<<< Updated upstream
	CreatePlayerWithCookie(id int, player *models.Player) error
	UpdatePlayerWithCookie(id int, player *models.Player) error
	ReadToken(id int) (*models.User, error)
	GetPlayerWithCookie(id int) (*models.Player, error)
=======
	UpdatePlayer(player *models.Player) error
	GetPlayerId(id int) (*models.Player, error)
>>>>>>> Stashed changes
=======
	CreatePlayer(player *models.Player) error
	UpdatePlayer(player *models.Player) error
	GetPlayerId(id int) (*models.Player, error)
>>>>>>> 281244cbd6c5e8c17cd2e03889eadb3996cf8ff1
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

<<<<<<< HEAD
<<<<<<< Updated upstream
func (p *playerRespository) CreatePlayerWithCookie(id int, player *models.Player) error {
=======
func (p *playerRespository) CreatePlayer(player *models.Player) error {
>>>>>>> 281244cbd6c5e8c17cd2e03889eadb3996cf8ff1

	if err := config.DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}
=======
// func (p *playerRespository) CreatePlayer(player *models.Player) error {

// 	if err := config.DB.Save(&player).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
>>>>>>> Stashed changes

func (p *playerRespository) UpdatePlayer(player *models.Player) error {

	if err := config.DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}

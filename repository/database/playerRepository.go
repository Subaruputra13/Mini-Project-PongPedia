package database

import (
	"PongPedia/config"
	"PongPedia/models"

	"gorm.io/gorm"
)

type PlayerRespository interface {
<<<<<<< Updated upstream
	CreatePlayerWithCookie(id int, player *models.Player) error
	UpdatePlayerWithCookie(id int, player *models.Player) error
	ReadToken(id int) (*models.User, error)
	GetPlayerWithCookie(id int) (*models.Player, error)
=======
	UpdatePlayer(player *models.Player) error
	GetPlayerId(id int) (*models.Player, error)
>>>>>>> Stashed changes
}

type playerRespository struct {
	db *gorm.DB
}

func NewPlayerRespository(db *gorm.DB) *playerRespository {
	return &playerRespository{db}
}

func (p *playerRespository) ReadToken(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *playerRespository) GetPlayerWithCookie(id int) (*models.Player, error) {
	var player models.Player

	if err := config.DB.Where("user_id = ?", id).Preload("Participation").First(&player).Error; err != nil {
		return nil, err
	}

	return &player, nil
}

<<<<<<< Updated upstream
func (p *playerRespository) CreatePlayerWithCookie(id int, player *models.Player) error {

	if err := config.DB.Model(&player).Where("user_id = ?", id).Save(&models.Player{
		Name:      player.Name,
		Age:       player.Age,
		BirthDate: player.BirthDate,
		Gender:    player.Gender,
		UserID:    id,
	}).Error; err != nil {
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

func (p *playerRespository) UpdatePlayerWithCookie(id int, player *models.Player) error {

	if err := config.DB.Model(&player).Where("user_id = ?", id).Updates(&models.Player{
		Name:      player.Name,
		Age:       player.Age,
		BirthDate: player.BirthDate,
		Gender:    player.Gender,
		UserID:    id,
	}).Error; err != nil {
		return err
	}
	return nil
}

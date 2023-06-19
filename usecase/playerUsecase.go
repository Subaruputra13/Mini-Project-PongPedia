package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"
	"time"
)

type PlayerUsecase interface {
	GetPlayers() ([]models.Player, error)
	GetPlayerByUserId(id uint) (player *models.Player, err error)
	UpdatePlayer(id uint, req *payload.CreateUpdatePlayerRequest) error
}

type playerUsecase struct {
	playerRespository database.PlayerRespository
	userRepository    database.UserRepository
}

func NewPlayerUsecase(
	playerRespository database.PlayerRespository,
	userRepository database.UserRepository,
) *playerUsecase {
	return &playerUsecase{playerRespository, userRepository}
}

func (p *playerUsecase) GetPlayers() (player []models.Player, err error) {
	player, err = p.playerRespository.GetPlayer()
	if err != nil {
		return nil, errors.New("Failed to get player")
	}

	return player, nil
}

func (p *playerUsecase) GetPlayerByUserId(id uint) (player *models.Player, err error) {
	player, err = p.playerRespository.GetPlayerUserId(id)
	if err != nil {
		return nil, errors.New("Failed to get player")
	}

	return player, nil
}

func (p *playerUsecase) UpdatePlayer(id uint, req *payload.CreateUpdatePlayerRequest) error {
	player, err := p.playerRespository.GetPlayerUserId(id)
	if err == nil {
		BirthDate, err := time.Parse("02/01/2006", req.BirthDate)
		if err != nil {
			return errors.New("Failed to parse birthdate")
		}

		player.Name = req.Name
		player.Age = req.Age
		player.BirthDate = &BirthDate
		player.Gender = req.Gender

		err = p.playerRespository.UpdatePlayer(player)
		if err != nil {
			return errors.New(err.Error())
		}
	} else {
		BirthDate, err := time.Parse("02/01/2006", req.BirthDate)
		if err != nil {
			return errors.New("Failed to parse birthdate")
		}

		userReq := &models.Player{
			Name:      req.Name,
			Age:       req.Age,
			BirthDate: &BirthDate,
			Gender:    req.Gender,
			UserID:    id,
		}

		err = p.playerRespository.CreatePlayer(userReq)
		if err != nil {
			return errors.New(err.Error())
		}
	}

	return nil
}

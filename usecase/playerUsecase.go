package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type PlayerUsecase interface {
	GetPlayer(id int) (*models.Player, error)
	CreatePlayer(id int, req *payload.CreateUpdatePlayerRequest) (res payload.PlayerResponse, err error)
	UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) (res payload.PlayerResponse, err error)
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

func (p *playerUsecase) GetPlayer(id int) (*models.Player, error) {

	user, err := p.userRepository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	player, err := p.playerRespository.GetPlayerId(int(user.ID))

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get player")
	}

	return player, nil
}

// Logic for Create and update player
func (p *playerUsecase) CreatePlayer(id int, req *payload.CreateUpdatePlayerRequest) (res payload.PlayerResponse, err error) {
	userReq := &models.Player{
		Name:      req.Name,
		Age:       req.Age,
		BirthDate: req.BirthDate,
		Gender:    req.Gender,
		UserID:    id,
	}

	err = p.playerRespository.CreatePlayer(userReq)

	if err != nil {
		echo.NewHTTPError(400, "Failed to create player")
		return
	}

	res = payload.PlayerResponse{
		ID:        int(userReq.ID),
		Name:      userReq.Name,
		Age:       userReq.Age,
		BirthDate: userReq.BirthDate,
		Gender:    userReq.Gender,
		UserID:    userReq.UserID,
	}

	return res, nil
}

func (p *playerUsecase) UpdatePlayer(id int, req *payload.CreateUpdatePlayerRequest) (res payload.PlayerResponse, err error) {

	player, err := p.playerRespository.GetPlayerId(id)

	player.Name = req.Name
	player.Age = req.Age
	player.BirthDate = req.BirthDate
	player.Gender = req.Gender

	if err != nil {
		echo.NewHTTPError(400, "Failed to get player")
		return
	}

	err = p.playerRespository.UpdatePlayer(player)

	if err != nil {
		echo.NewHTTPError(400, "Failed to update player")
		return
	}

	res = payload.PlayerResponse{
		ID:        int(player.ID),
		Name:      player.Name,
		Age:       player.Age,
		BirthDate: player.BirthDate,
		Gender:    player.Gender,
		UserID:    player.UserID,
	}

	return res, nil
}

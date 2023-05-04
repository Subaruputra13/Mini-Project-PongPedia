package usecase

import (
	"PongPedia/models"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type PlayerUsecase interface {
	CreatePlayer(id int, player *models.Player) (*models.User, error)
	UpdatePlayer(id int, player *models.Player) (*models.User, error)
	GetPlayer(id int) (*models.Player, error)
}

type playerUsecase struct {
	playerRespository database.PlayerRespository
}

func NewPlayerUsecase(playerRespository database.PlayerRespository) *playerUsecase {
	return &playerUsecase{playerRespository}
}

func (p *playerUsecase) GetPlayer(id int) (*models.Player, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	player, err := p.playerRespository.GetPlayerWithCookie(int(user.ID))

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get player")
	}

	return player, nil
}

// Logic for update player
func (p *playerUsecase) CreatePlayer(id int, player *models.Player) (*models.User, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	if err := p.playerRespository.CreatePlayerWithCookie(id, player); err != nil {
		return nil, echo.NewHTTPError(400, "Failed to create player")
	}

	return user, nil
}

// Logic for update player
func (p *playerUsecase) UpdatePlayer(id int, player *models.Player) (*models.User, error) {

	user, err := p.playerRespository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to read token")
	}

	if err := p.playerRespository.UpdatePlayerWithCookie(id, player); err != nil {
		return nil, echo.NewHTTPError(400, "Failed to create player")
	}

	return user, nil
}

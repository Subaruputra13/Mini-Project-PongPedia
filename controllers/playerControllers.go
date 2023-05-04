package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type PlayerController interface {
	CreatePlayerController(c echo.Context) error
	UpdatePlayerController(c echo.Context) error
}

type playerController struct {
	playerUsecase     usecase.PlayerUsecase
	playerRespository database.PlayerRespository
}

func NewPlayerController(
	playerUsecase usecase.PlayerUsecase,
	playerRespository database.PlayerRespository,
) *playerController {
	return &playerController{playerUsecase, playerRespository}
}

func (p *playerController) GetPlayerController(c echo.Context) error {
	id := m.Auth(c)

	player, err := p.playerUsecase.GetPlayer(id)

	if err != nil {
		return echo.NewHTTPError(400, "Failed to read token")
	}

	return c.JSON(200, payload.Response{
		Message: "Success get user",
		Data:    player,
	})
}

func (p *playerController) CreatePlayerController(c echo.Context) error {
	request := payload.CreateUpdatePlayerRequest{}

	c.Bind(&request)

	id := m.Auth(c)

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	player := models.Player{
		Name:      request.Name,
		Age:       request.Age,
		BirthDate: request.BirthDate,
		Gender:    request.Gender,
		UserID:    id,
	}

	if _, err := p.playerUsecase.CreatePlayer(id, &player); err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	playerResponse := payload.PlayerResponse{
		Name:      player.Name,
		Age:       player.Age,
		BirthDate: player.BirthDate,
		Gender:    player.Gender,
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    playerResponse,
	})
}

// Controllers for player
func (p *playerController) UpdatePlayerController(c echo.Context) error {
	request := payload.CreateUpdatePlayerRequest{}

	c.Bind(&request)

	id := m.Auth(c)

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	player := models.Player{
		Name:      request.Name,
		Age:       request.Age,
		BirthDate: request.BirthDate,
		Gender:    request.Gender,
	}

	if _, err := p.playerUsecase.UpdatePlayer(id, &player); err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	playerResponse := payload.PlayerResponse{
		Name:      player.Name,
		Age:       player.Age,
		BirthDate: player.BirthDate,
		Gender:    player.Gender,
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    playerResponse,
	})
}

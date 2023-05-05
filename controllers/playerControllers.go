package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type PlayerController interface {
	GetPlayerController(c echo.Context) error
	CreatePlayerController(c echo.Context) error
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
	id, _ := m.IsUser(c)

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
	req := payload.CreateUpdatePlayerRequest{}

	c.Bind(&req)

	id, _ := m.IsUser(c)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := p.playerUsecase.CreatePlayer(id, &req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(201, payload.Response{
		Message: "Success update user",
		Data:    res,
	})
}

func (p *playerController) UpdatePlayerController(c echo.Context) error {
	req := payload.CreateUpdatePlayerRequest{}

	c.Bind(&req)

	id, _ := m.IsUser(c)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := p.playerUsecase.UpdatePlayer(id, &req)

	if err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    res,
	})
}

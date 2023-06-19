package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type PlayerController interface {
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

func (p *playerController) GetPlayerByUserIdController(c echo.Context) error {
	userId, err := m.IsUser(c)
	if err != nil {
		return c.JSON(400, "this routes only for user")
	}

	player, err := p.playerUsecase.GetPlayerByUserId(uint(userId))
	if err != nil {
		return c.JSON(400, "Failed to get player")
	}

	return c.JSON(200, player)
}

func (p *playerController) UpdatePlayerController(c echo.Context) error {
	req := payload.CreateUpdatePlayerRequest{}

	c.Bind(&req)

	userId, err := m.IsUser(c)
	if err != nil {
		return c.JSON(400, "this routes only for user")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	err = p.playerUsecase.UpdatePlayer(uint(userId), &req)
	if err != nil {
		return echo.NewHTTPError(400, "Username already exist")
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Success update user",
	})
}

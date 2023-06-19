package controllers

import (
	"PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type AdminControllers interface {
	DashboardAdminController(c echo.Context) error
	GetAllUserController(c echo.Context) error
	CreateMatchController(c echo.Context) error
	UpdateMatchController(c echo.Context) error
}

type adminControllers struct {
	adminUsecase     usecase.DashboardUsecase
	matchUsecase     usecase.MatchUsecase
	playerUsecase    usecase.PlayerUsecase
	turnamentUsecase usecase.TurnamentUsecase
}

func NewAdminControllers(
	adminUsecase usecase.DashboardUsecase,
	matchUsecase usecase.MatchUsecase,
	playerUsecase usecase.PlayerUsecase,
	turnamentUsecase usecase.TurnamentUsecase,
) *adminControllers {
	return &adminControllers{adminUsecase, matchUsecase, playerUsecase, turnamentUsecase}
}

func (a *adminControllers) DashboardAdminController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	result := a.adminUsecase.DashboardAdmin()

	return c.JSON(200, payload.Response{
		Message: "success get admin dashboard",
		Data:    result,
	})
}

func (a *adminControllers) GetAllUserController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	user, err := a.adminUsecase.GetAllUser()
	if err != nil {
		return c.JSON(400, "failed get all user")
	}

	return c.JSON(200, payload.Response{
		Message: "success get all user",
		Data:    user,
	})
}

func (a *adminControllers) CreateMatchController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	req := payload.CreateMatchRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := a.matchUsecase.CreateMatch(&req)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Match",
		Data:    res,
	})
}

func (a *adminControllers) UpdateMatchController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	match, err := a.matchUsecase.GetMatchById(uint(id))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	c.Bind(&match)

	err = a.matchUsecase.UpdateMatch(match)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "Success Update Match")
}

func (a *adminControllers) GetAllPlayersController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	player, err := a.playerUsecase.GetPlayers()
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get user",
		Data:    player,
	})
}

func (a *adminControllers) CreateTurnamentController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(401, "this routes is for admin only")
	}

	req := payload.TurnamentRequest{}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return c.JSON(400, "Field cannot be empty")
	}

	turnament, err := a.turnamentUsecase.CreateTurnament(&req)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success create turnament",
		Data:    turnament,
	})
}

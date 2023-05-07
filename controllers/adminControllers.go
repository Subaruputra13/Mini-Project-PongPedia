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
	adminUsecase usecase.DashboardUsecase
	matchUsecase usecase.MatchUsecase
}

func NewAdminControllers(
	adminUsecase usecase.DashboardUsecase,
	matchUsecase usecase.MatchUsecase,
) *adminControllers {
	return &adminControllers{adminUsecase, matchUsecase}
}

func (a *adminControllers) DashboardAdminController(c echo.Context) error {

	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	result := a.adminUsecase.DashboardAdmin()

	return c.JSON(200, payload.Response{
		Message: "success get admin dashboard",
		Data:    result,
	})
}

func (a *adminControllers) GetAllUserController(c echo.Context) error {

	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
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
	req := payload.CreateMatchRequest{}

	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	err := a.matchUsecase.CreateMatch(&req)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Create Match",
	})
}

func (a *adminControllers) UpdateMatchController(c echo.Context) error {
	req := payload.UpdateMatchRequest{}

	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(400, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	err := a.matchUsecase.UpdateMatch(&req, id)
	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, "Success Update Match")
}

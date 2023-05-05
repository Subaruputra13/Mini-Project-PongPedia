package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"strconv"

	"github.com/labstack/echo"
)

type TurnamentControllers interface{}

type turnamentControllers struct {
	turnamanetUsecase   usecase.TurnamentUsecase
	turnamentRepository database.TurnamentRepository
}

func NewTurnamentControllers(
	turnamanetUsecase usecase.TurnamentUsecase,
	turnamentRepository database.TurnamentRepository,
) *turnamentControllers {
	return &turnamentControllers{turnamanetUsecase, turnamentRepository}
}

func (t *turnamentControllers) GetTurnamentController(c echo.Context) error {
	turnament, err := t.turnamanetUsecase.GetTurnament()

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get turnament",
		Data:    turnament,
	})
}

func (t *turnamentControllers) GetTurnamentDetailController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	turnament, err := t.turnamanetUsecase.GetTurnamentById(id)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success get turnament",
		Data:    turnament,
	})
}

func (t *turnamentControllers) CreateTurnamentController(c echo.Context) error {
	req := payload.TurnamentRequest{}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid input")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	turnament, err := t.turnamanetUsecase.CreateTurnament(&req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success create turnament",
		Data:    turnament,
	})
}

func (t *turnamentControllers) RegisterTurnamentController(c echo.Context) error {
	req := payload.RegisterTurnamentRequest{}

	c.Bind(&req)

	id, _ := m.IsUser(c)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	err := t.turnamanetUsecase.RegisterTurnament(id, &req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, "Success register turnament")
}

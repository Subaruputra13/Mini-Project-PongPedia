package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"
	"fmt"

	"github.com/labstack/echo"
)

type UserController interface {
	GetUserController(c echo.Context) error
	UpdateUserController(c echo.Context) error
	CreatePlayerController(c echo.Context) error
	UpdatePlayerController(c echo.Context) error
}

type userController struct {
	userUsecase    usecase.UserUsecase
	userRepository database.UserRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
) *userController {
	return &userController{userUsecase, userRepository}
}

func (u *userController) GetUserController(c echo.Context) error {

	userId, err := m.IsUser(c)
	if err != nil {
		return c.JSON(401, "this routes is for user only")
	}

	res, err := u.userUsecase.GetUserById(uint(userId))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: fmt.Sprintf("Welcome %s", res.Username),
		Data:    res,
	})
}

func (u *userController) UpdateUserController(c echo.Context) error {
	userId, err := m.IsUser(c)
	if err != nil {
		return c.JSON(401, "this routes is for user only")
	}

	payloadUser := payload.UpdateUserRequest{}

	c.Bind(&payloadUser)

	user, err := u.userUsecase.GetUserById(uint(userId))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	c.Bind(user)

	res, err := u.userUsecase.UpdateUser(user, &payloadUser)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    res,
	})
}

func (u *userController) DeleteUserController(c echo.Context) error {
	userId, err := m.IsUser(c)
	if err != nil {
		return c.JSON(401, "this routes is for user only")
	}

	payloadUser := payload.DeleteUserRequest{}

	c.Bind(&payloadUser)

	user, err := u.userUsecase.GetUserById(uint(userId))
	if err != nil {
		return c.JSON(400, err.Error())
	}

	c.Bind(user)

	if err := u.userUsecase.DeleteUser(user, &payloadUser); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, "Succes Delete User")
}

package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models"
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

	id := m.Auth(c)

	user, err := u.userUsecase.GetUserById(id)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: fmt.Sprintf("Welcome %s", user.Username),
		Data:    user,
	})
}

func (u *userController) UpdateUserController(c echo.Context) error {
	request := payload.UpdateUserRequest{}

	id := m.Auth(c)

	c.Bind(&request)

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := u.userUsecase.UpdateUser(id, &user); err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    user,
	})
}

func (u *userController) DeleteUserController(c echo.Context) error {
	id := m.Auth(c)

	password := c.FormValue("Password")

	if _, err := u.userRepository.DeleteUser(id, password); err != nil {
		return echo.NewHTTPError(400, "Wrong Password !")
	}

	m.DeleteCookie(c)

	return c.JSON(200, "Succes Delete User")
}

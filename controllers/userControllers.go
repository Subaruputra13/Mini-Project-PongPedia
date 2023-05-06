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

	id, _ := m.IsUser(c)

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
	req := payload.UpdateUserRequest{}

	id, _ := m.IsUser(c)

<<<<<<< HEAD
<<<<<<< Updated upstream
	c.Bind(&request)
=======
	c.Bind(&req)
>>>>>>> Stashed changes
=======
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Failed to bind user")
	}
>>>>>>> 281244cbd6c5e8c17cd2e03889eadb3996cf8ff1

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := u.userUsecase.UpdateUser(id, &req)
	if err != nil {
		return echo.NewHTTPError(400, "failed to update user")
	}

	return c.JSON(200, payload.Response{
		Message: "Success update user",
		Data:    res,
	})
}

func (u *userController) DeleteUserController(c echo.Context) error {
	id, _ := m.IsUser(c)

	password := c.FormValue("Password")

	if err := u.userUsecase.DeleteUser(id, password); err != nil {
		return echo.NewHTTPError(400, "Wrong Password !")
	}

	return c.JSON(200, "Succes Delete User")
}

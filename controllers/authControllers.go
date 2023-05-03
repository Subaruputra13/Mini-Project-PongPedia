package controllers

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type UserController interface {
	LoginUserController(c echo.Context) error
	RegisterUserController(c echo.Context) error
}

type userControler struct {
	userUsacase    usecase.UserUsecase
	userRepository database.UserRepository
}

func NewUserController(
	userUsecase usecase.UserUsecase,
	userRepository database.UserRepository,
) *userControler {
	return &userControler{userUsecase, userRepository}
}

func (u *userControler) LoginUserController(c echo.Context) error {
	request := payload.LoginRequest{}

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	user, err := u.userUsacase.LoginUser(request.Email, request.Password, "PLAYER")

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userResponse := payload.LoginResponse{Email: user.Email, Token: user.Token}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    userResponse,
	})
}

func (u *userControler) RegisterUserController(c echo.Context) error {
	request := payload.RegisterRequest{}

	c.Bind(&request)

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	if _, err := u.userRepository.GetuserByEmail(request.Email); err == nil {
		return echo.NewHTTPError(400, "Email Already Registered")
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}

	if err := u.userUsacase.CreateUser(&user); err != nil {
		return echo.NewHTTPError(400, "Failed to Register User")
	}

	userResponse := payload.RegisterResponse{Username: user.Username, Email: user.Email, Password: user.Password, Role: user.Role}

	return c.JSON(200, payload.Response{
		Message: "Success Register User",
		Data:    userResponse,
	})

}

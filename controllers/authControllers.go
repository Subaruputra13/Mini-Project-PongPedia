package controllers

import (
	m "PongPedia/middleware"
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
	RegisterUserController(c echo.Context) error
}

type authControler struct {
	authUsecase    usecase.AuthUsecase
	authRepository database.AuthRepository
	userRepository database.UserRepository
}

func NewAuthController(
	authUsecase usecase.AuthUsecase,
	authRepository database.AuthRepository,
	userRepository database.UserRepository,
) *authControler {
	return &authControler{
		authUsecase,
		authRepository,
		userRepository,
	}
}

func (a *authControler) LoginUserController(c echo.Context) error {
	request := payload.LoginRequest{}

<<<<<<< Updated upstream
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := a.authUsecase.LoginUser(&user); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
=======
	c.Bind(&req)
>>>>>>> Stashed changes

	m.CreateCookie(c, user.Token)

	userResponse := payload.LoginResponse{Email: user.Email, Token: user.Token}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    userResponse,
	})
}

func (a *authControler) RegisterUserController(c echo.Context) error {
	request := payload.RegisterRequest{}

	c.Bind(&request)

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}

	if err := a.userRepository.CreateUser(&user); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userResponse := payload.RegisterResponse{Username: user.Username, Email: user.Email, Password: user.Password, Role: user.Role}

	return c.JSON(200, payload.Response{
		Message: "Success Register User",
		Data:    userResponse,
	})

}

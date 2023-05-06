package controllers

import (
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
	userUsecase    usecase.UserUsecase
}

func NewAuthController(
	authUsecase usecase.AuthUsecase,
	authRepository database.AuthRepository,
	userUsecase usecase.UserUsecase,
) *authControler {
	return &authControler{
		authUsecase,
		authRepository,
		userUsecase,
	}
}

func (a *authControler) LoginUserController(c echo.Context) error {
	req := payload.LoginRequest{}

<<<<<<< HEAD
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
=======
	if err := c.Bind(&req); err != nil {
>>>>>>> 281244cbd6c5e8c17cd2e03889eadb3996cf8ff1
		return echo.NewHTTPError(400, err.Error())
	}
=======
	c.Bind(&req)
>>>>>>> Stashed changes

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Field cannot be empty")
	}

	res, err := a.authUsecase.LoginUser(&req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    res,
	})
}

func (a *authControler) RegisterUserController(c echo.Context) error {
	req := payload.RegisterRequest{}

	c.Bind(&req)

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid Request")
	}

	res, err := a.userUsecase.CreateUser(&req)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Register User",
		Data:    res,
	})

}

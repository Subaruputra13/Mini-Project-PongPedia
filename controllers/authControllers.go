package controllers

import (
	"PongPedia/models/payload"
	"PongPedia/usecase"

	"github.com/labstack/echo"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
	RegisterUserController(c echo.Context) error
}

type authControler struct {
	authUsecase usecase.AuthUsecase
	userUsecase usecase.UserUsecase
}

func NewAuthController(
	authUsecase usecase.AuthUsecase,
	userUsecase usecase.UserUsecase,
) *authControler {
	return &authControler{
		authUsecase,
		userUsecase,
	}
}

func (a *authControler) LoginUserController(c echo.Context) error {
	req := payload.LoginRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error payload login",
			"error":   err.Error(),
		})
	}

	res, err := a.authUsecase.LoginUser(&req)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    res,
	})
}

func (a *authControler) RegisterUserController(c echo.Context) error {
	req := payload.RegisterRequest{}

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error payload login",
			"error":   err.Error(),
		})
	}

	err := a.userUsecase.CreateUser(&req)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Success Register",
	})

}

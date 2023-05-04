package usecase

import (
	"PongPedia/middleware"
	"PongPedia/models"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type AuthUsecase interface {
	LoginUser(user *models.User) error
}

type authUsecase struct {
	authRepository database.AuthRepository
}

func NewAuthUsecase(authRepository database.AuthRepository) *authUsecase {
	return &authUsecase{authRepository}
}

// Logic for login user
func (a *authUsecase) LoginUser(user *models.User) error {
	if _, err := a.authRepository.CheckEmail(user.Email); err != nil {
		return echo.NewHTTPError(400, "Email not registered")
	}

	if err := a.authRepository.LoginUser(user); err != nil {
		return echo.NewHTTPError(400, "Wrong password")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)

	if err != nil {
		return echo.NewHTTPError(400, "Failed to generate token")
	}

	user.Token = token

	return nil
}

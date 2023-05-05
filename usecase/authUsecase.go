package usecase

import (
	"PongPedia/middleware"
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type AuthUsecase interface {
	LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error)
}

type authUsecase struct {
	authRepository database.AuthRepository
	userRepository database.UserRepository
}

func NewAuthUsecase(
	authRepository database.AuthRepository,
	userRepository database.UserRepository,
) *authUsecase {
	return &authUsecase{authRepository, userRepository}
}

// Logic for login user
func (a *authUsecase) LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error) {
	userReq := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	_, err = a.userRepository.GetuserByEmail(userReq.Email)

	if err != nil {
		echo.NewHTTPError(400, "Email not registered")
		return
	}

	err = a.authRepository.LoginUser(userReq)

	if err != nil {
		echo.NewHTTPError(400, "Wrong password")
		return
	}

	token, err := middleware.CreateToken(int(userReq.ID), userReq.Role)

	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	userReq.Token = token

	res = payload.LoginResponse{
		Email: userReq.Email,
		Token: userReq.Token,
	}

	return
}

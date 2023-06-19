package usecase

import (
	"PongPedia/middleware"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error)
}

type authUsecase struct {
	userRepository database.UserRepository
}

func NewAuthUsecase(
	userRepository database.UserRepository,
) *authUsecase {
	return &authUsecase{userRepository}
}

// Logic for login user
func (a *authUsecase) LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error) {

	user, err := a.userRepository.GetuserByEmail(req.EmailOrUsername, req.EmailOrUsername)
	if err != nil {
		return res, errors.New("User not already registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New("Password is incorrect")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return res, errors.New("Failed to create token")
	}

	user.Token = token

	res = payload.LoginResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}

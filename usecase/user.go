package usecase

import (
	"PongPedia/middleware"
	"PongPedia/models"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type UserUsecase interface {
	CreateUser(user *models.User) error
	GetuserByEmail(email string) (*models.User, error)
	LoginUser(email, password, role string) (*models.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) CreateUser(user *models.User) error {
	if err := u.userRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GetuserByEmail(email string) (*models.User, error) {
	return u.userRepository.GetuserByEmail(email)
}

func (u *userUsecase) LoginUser(email, password, role string) (*models.User, error) {
	user, err := u.userRepository.GetuserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(400, "Email not Registered")
	}

	if user.Password != password {
		return nil, echo.NewHTTPError(400, "Password Incorrect")
	}

	if user.Role != role {
		return nil, echo.NewHTTPError(401, "Unauthorized")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to generate token")
	}

	user.Token = token

	return user, nil
}

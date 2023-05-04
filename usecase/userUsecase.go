package usecase

import (
	"PongPedia/models"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type UserUsecase interface {
	GetUserById(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(id int, password string) (user *models.User, err error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepository database.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

// Logic for get user with cookie
func (u *userUsecase) GetUserById(id int) (*models.User, error) {
	user, err := u.userRepository.GetUseByIdWithCookie(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get user")
	}

	return user, nil
}

// Logic for create user
func (u *userUsecase) CreateUser(user *models.User) error {
	if err := u.userRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

// Logic for update user
func (u *userUsecase) UpdateUser(id int, user *models.User) error {

	if err := u.userRepository.UpdateUserWithCookie(id, user); err != nil {
		return err
	}

	return nil
}

// Logic for Delete user
func (u *userUsecase) DeleteUser(id int, password string) (*models.User, error) {

	user, err := u.userRepository.ReadToken(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}

	_, err = u.userRepository.DeleteUser(id, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

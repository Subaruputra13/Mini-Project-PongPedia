package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"

	"github.com/labstack/echo"
)

type UserUsecase interface {
	GetUserById(id int) (*models.User, error)
	UpdateUser(id int, req *payload.UpdateUserRequest) (res payload.UpdateUserRequest, err error)
	CreateUser(reqs *payload.RegisterRequest) (res payload.RegisterResponse, err error)
	DeleteUser(id int, password string) error
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepository database.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

// Logic for get user with cookie
func (u *userUsecase) GetUserById(id int) (*models.User, error) {
	user, err := u.userRepository.GetUseById(id)

	if err != nil {
		return nil, echo.NewHTTPError(400, "Failed to get user")
	}

	return user, nil
}

// Logic for update user
func (u *userUsecase) UpdateUser(id int, req *payload.UpdateUserRequest) (res payload.UpdateUserRequest, err error) {
	userReq := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	userReq.ID = uint(id)

	err = u.userRepository.UpdateUser(userReq)

	if err != nil {
		echo.NewHTTPError(400, "Failed to update user")
		return
	}

	res = payload.UpdateUserRequest{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	return
}

// Logic for create user
func (u *userUsecase) CreateUser(reqs *payload.RegisterRequest) (res payload.RegisterResponse, err error) {

	userReq := &models.User{
		Username: reqs.Username,
		Email:    reqs.Email,
		Password: reqs.Password,
	}

	err = u.userRepository.CreateUser(userReq)
	if err != nil {
		echo.NewHTTPError(400, "Failed to create user")
		return
	}

	res = payload.RegisterResponse{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
		Role:     userReq.Role,
	}

	return
}

// Logic for Delete user
func (u *userUsecase) DeleteUser(id int, password string) error {

	user, err := u.userRepository.ReadToken(id)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	if user.Password != password {
		return echo.NewHTTPError(400, "Wrong Password")
	}

	err = u.userRepository.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}

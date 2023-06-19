package usecase

import (
	"PongPedia/models"
	"PongPedia/models/payload"
	"PongPedia/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUserById(id uint) (user *models.User, err error)
	CreateUser(reqs *payload.RegisterRequest) error
	UpdateUser(user *models.User, req *payload.UpdateUserRequest) (res payload.GetAllUserResponse, err error)
	DeleteUser(user *models.User, req *payload.DeleteUserRequest) error
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepository database.UserRepository) *userUsecase {
	return &userUsecase{userRepository}
}

// Logic for get user with cookie
func (u *userUsecase) GetUserById(id uint) (user *models.User, err error) {
	user, err = u.userRepository.GetUserById(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

// Logic for create user
func (u *userUsecase) CreateUser(reqs *payload.RegisterRequest) error {

	userReq := &models.User{
		Username: reqs.Username,
		Email:    reqs.Email,
		Password: reqs.Password,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Failed to hash password")
	}

	userReq.Password = string(passwordHash)

	err = u.userRepository.CreateUser(userReq)
	if err != nil {
		return errors.New("Failed to create user")
	}

	return nil
}

// Logic for update user
func (u *userUsecase) UpdateUser(user *models.User, req *payload.UpdateUserRequest) (res payload.GetAllUserResponse, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, errors.New("Failed to hash password")
	}

	user.Password = string(passwordHash)

	err = u.userRepository.UpdateUser(user)
	if err != nil {
		return res, errors.New("Failed to update user")
	}

	res = payload.GetAllUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return res, nil
}

// Logic for Delete user
func (u *userUsecase) DeleteUser(user *models.User, req *payload.DeleteUserRequest) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.New("Password is incorrect")
	}

	err = u.userRepository.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}

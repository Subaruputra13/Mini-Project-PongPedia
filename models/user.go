package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'; not-null"`
}

type UserReponse struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type CustomValidator struct {
	Validators *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validators.Struct(i)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

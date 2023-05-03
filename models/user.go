package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required"`
	Password string `json:"password" form:"password" gorm:"unique" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'; not-null"`
	Player   Player
}

// Only for response Token
type UserReponse struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

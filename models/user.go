package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password" gorm:"unique"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'; not-null"`
	Player   Player `json:"player" gorm:"foreignKey:UserID"`
}

// Only for response Token
type UserReponse struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

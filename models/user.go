package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'"`
	Token    string `json:"-" gorm:"-"`
	Player   Player `gorm:"foreignKey:UserID"`
}

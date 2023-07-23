package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique;not null"`
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin" gorm:"default:false"`
	Token    string `json:"-" gorm:"-"`
	Player   *Player
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Players  Player `gorm:"foreignKey:UserId"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'; not-null"`
}

type UserReponse struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

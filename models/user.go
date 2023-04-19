package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:enum('ADMIN', 'PLAYER');default:'PLAYER'; not-null"`
	Player   Player `gorm:"foreignKey:UserId"`
}

type UserReponse struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

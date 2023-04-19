package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	FullName string `json:"fullname" form:"fullname"`
	Age      int    `json:"age" form:"age"`
	Domisili string `json:"domisili" form:"domisili"`
	UserId   int    `gorm:"unique" json:"user_id" form:"user_id"`
}

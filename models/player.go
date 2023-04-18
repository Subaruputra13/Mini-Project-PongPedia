package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
	Domisili string `json:"domisili" form:"domisili"`
	UserId   int    `json:"-" form:"user_id"`
}

package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	FullName    string        `json:"fullname" form:"fullname"`
	Age         int           `json:"age" form:"age"`
	Domisili    string        `json:"domisili" form:"domisili"`
	UserID      uint          `json:"user_id" form:"user_id"`
	User        UserResponses `json:"user"`
	Turnament   []Turnament   `gorm:"many2many:matches" json:"turnament"`
	TurnamentId []int         `json:"turnament_id" form:"turnament_id" gorm:"-"`
}

type PlayerResponse struct {
	gorm.Model
	FullName string `json:"fullname" form:"fullname"`
	Age      uint   `json:"age" form:"age"`
	Domisili string `json:"domisili" form:"domisili"`
	UserID   uint   `json:"-" form:"user_id"`
}

func (PlayerResponse) TableName() string {
	return "players"
}

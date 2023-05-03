package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name          string          `json:"name" form:"name"`
	Age           int             `json:"age" form:"age"`
	BirthDate     string          `json:"birth_date" form:"birth_date"`
	Gender        string          `json:"gender" form:"gender"`
	UserID        int             `json:"user_id" form:"user_id"`
	Participation []Participation `gorm:"foreignKey:PlayerID"`
}

type PlayerResponse struct {
	ID        uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	Age       int    `json:"age" form:"age"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Gender    string `json:"gender" form:"gender"`
	UserID    int    `json:"-" form:"user_id"`
}

func (PlayerResponse) TableName() string {
	return "players"
}

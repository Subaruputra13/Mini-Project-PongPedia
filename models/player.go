package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name          string     `json:"name" form:"name" gorm:"unique"`
	Age           uint       `json:"age" form:"age"`
	BirthDate     *time.Time `json:"birth_date" form:"birth_date"`
	Gender        string     `json:"gender" form:"gender" gorm:"type:enum('Male', 'Female')"`
	UserID        uint       `json:"user_id" form:"user_id" gorm:"unique"`
	Participation []Participation
}

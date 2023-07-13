package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	FirstName     string     `json:"first_name" form:"first_name"`
	LastName      string     `json:"last_name" form:"last_name"`
	Age           uint       `json:"age" form:"age"`
	BirthDate     *time.Time `json:"birth_date" form:"birth_date"`
	Gender        string     `json:"gender" form:"gender" gorm:"type:enum('Male', 'Female')"`
	Styles        string     `json:"styles" form:"styles" gorm:"type:enum('Right Hand', 'Left Hand')"`
	UserID        uint       `json:"user_id" form:"user_id" gorm:"unique"`
	Participation []Participation
}

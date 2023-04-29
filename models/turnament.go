package models

import (
	"gorm.io/gorm"
)

type Turnament struct {
	gorm.Model
	Name      string  `json:"name" form:"name"`
	StartDate string  `json:"start_date" form:"start_date"`
	EndDate   string  `json:"end_date" form:"end_date"`
	Location  string  `json:"location" form:"location"`
	Matchs    []Match `gorm:"foreignKey:TurnamentID"`
}

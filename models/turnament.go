package models

import (
	"gorm.io/gorm"
)

type Turnament struct {
	gorm.Model
	Name          string `json:"name" form:"name" gorm:"unique"`
	StartDate     string `json:"start_date" form:"start_date"`
	EndDate       string `json:"end_date" form:"end_date"`
	Location      string `json:"location" form:"location"`
	Slot          int    `json:"slot" form:"slot"`
	Participation []Participation
	Match         []Match
}

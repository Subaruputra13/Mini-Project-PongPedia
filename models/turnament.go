package models

import (
	"time"

	"gorm.io/gorm"
)

type Turnament struct {
	gorm.Model
	Name          string     `json:"name" form:"name"`
	StartDate     *time.Time `json:"start_date" form:"start_date"`
	EndDate       *time.Time `json:"end_date" form:"end_date"`
	Location      string     `json:"location" form:"location"`
	Slot          uint       `json:"slot" form:"slot"`
	Participation []Participation
	Match         []Match
}

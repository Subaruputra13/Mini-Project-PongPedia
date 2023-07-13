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
	Place         string     `json:"place" form:"place"`
	Address       string     `json:"address" form:"address"`
	Village       string     `json:"village" form:"village"`
	Prize         uint       `json:"prize" form:"prize"`
	Slot          uint       `json:"slot" form:"slot"`
	Champion      string     `json:"champion" form:"champion"`
	Participation []Participation
	Match         []Match
}

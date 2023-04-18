package models

import (
	"time"

	"gorm.io/gorm"
)

type Turnamnet struct {
	gorm.Model
	Name         string     `json:"name" form:"name"`
	Date         *time.Time `json:"date" form:"date"`
	TunamentType string     `json:"turnament_type" form:"turnament_type"`
	Matchs       []Match    `gorm:"foreignKey:TurnamentId"`
}

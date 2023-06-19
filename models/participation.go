package models

import "gorm.io/gorm"

type Participation struct {
	gorm.Model
	PlayerID    uint `json:"player_id" form:"player_id"`
	TurnamentID uint `json:"turnament_id" form:"turnament_id"`
	Player      *Player
	Turnament   *Turnament
}

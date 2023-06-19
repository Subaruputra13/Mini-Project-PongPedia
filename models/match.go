package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	MatchName      string     `json:"match_name" form:"match_name"`
	MatchDate      *time.Time `json:"match_date" form:"match_date"`
	Player_1       uint       `json:"player_1" form:"player_1"`
	Player_2       uint       `json:"player_2" form:"player_2"`
	Player_1_Score uint       `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint       `json:"player_2_score" form:"player_2_score"`
	TurnamentID    uint       `json:"turnament_id" form:"turnament_id"`
	Turnament      *Turnament
}

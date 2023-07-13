package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	MatchType      string     `json:"match_type" form:"match_type" gorm:"type:enum('Round of 16', 'Quarter Final', 'Semi Final', 'Final')"`
	MatchDate      *time.Time `json:"match_date" form:"match_date"`
	Player_1ID     uint       `json:"player_1ID" form:"player_1ID"`
	Player_1       *Player
	Player_2ID     uint `json:"player_2ID" form:"player_2ID"`
	Player_2       *Player
	Player_1_Score uint `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint `json:"player_2_score" form:"player_2_score"`
	TurnamentID    uint `json:"turnament_id" form:"turnament_id"`
	Turnament      *Turnament
}

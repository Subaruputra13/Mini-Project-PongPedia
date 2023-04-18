package models

import "time"

type Match struct {
	Date        *time.Time `json:"date" form:"date"`
	Result      string     `json:"result" form:"result"`
	PlayerId    int        `json:"player_id" form:"player_id"`
	TurnamentId int        `json:"turnament_id" form:"turnament_id"`
}

package models

type Match struct {
	PlayerId    int `json:"player_id" form:"player_id"`
	TurnamentId int `json:"turnament_id" form:"turnament_id"`
}

package models

type Turnament struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Date          string `json:"date" form:"date"`
	TurnamentType string `json:"turnament_type" form:"turnament_type"`
}

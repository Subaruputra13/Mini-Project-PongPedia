package models

type MatchDetails struct {
	ID          uint   `json:"id" form:"id"`
	Date        string `json:"date" form:"date"`
	Result      string `json:"result" form:"result"`
	PlayerId    int    `gorm:"primarykey"`
	TurnamentId int    `gorm:"primarykey"`
}

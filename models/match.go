package models

type Match struct {
	ID          int    `json:"id" form:"id" gorm:"primary_key"`
	PlayerID    int    `json:"player_id" form:"player_id"`
	TurnamentID int    `json:"turnament_id" form:"turnament_id"`
	MatchDate   string `json:"match_date" form:"match_date"`
	Result      string `json:"result" form:"result"`
}

type MatchResponse struct {
	ID           int    `json:"id" form:"id" gorm:"primary_key"`
	PlayerID     int    `json:"player_id" form:"player_id"`
	TurnamentID  int    `json:"turnament_id" form:"turnament_id"`
	MatchDate    string `json:"match_date" form:"match_date"`
	Player1Score int    `json:"player1_score" form:"player1_score"`
	Player2Score int    `json:"player2_score" form:"player2_score"`
	Result       string `json:"result" form:"result"`
}

func (MatchResponse) TableName() string {
	return "Matchs"
}

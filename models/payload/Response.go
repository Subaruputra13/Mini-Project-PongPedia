package payload

import (
	"PongPedia/models"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetAllUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Player   *models.Player
}

type PlayerResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	UserID    int    `json:"-"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type TurnamentDetailResponse struct {
	Name          string `json:"name" form:"name"`
	StartDate     string `json:"start_date" form:"start_date"`
	EndDate       string `json:"end_date" form:"end_date"`
	Location      string `json:"location" form:"location"`
	Paticipations []ParticipationResponse
}
type TurnamentResponse struct {
	ID        uint       `json:"id" form:"id"`
	Name      string     `json:"name" form:"name"`
	StartDate *time.Time `json:"start_date" form:"start_date"`
	EndDate   *time.Time `json:"end_date" form:"end_date"`
	Location  string     `json:"location" form:"location"`
	Slot      uint       `json:"slot" form:"slot"`
}

type ParticipationResponse struct {
	PlayerID int `json:"player_id"`
	Player   PlayerResponse
}

type DashboardAdminResponse struct {
	TotalUser      int64 `json:"total_user"`
	TotalPlayer    int64 `json:"total_player"`
	TotalTurnament int64 `json:"total_turnament"`
	TotalMatch     int64 `json:"total_match"`
}

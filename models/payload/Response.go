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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Style     string `json:"style"`
	UserID    int    `json:"-"`
}

type LoginResponse struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
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
	Place         string `json:"place" `
	Address       string `json:"address" `
	Village       string `json:"village" `
	Prize         uint   `json:"prize" `
	Slot          uint   `json:"slot" `
	Champion      string `json:"champion" `
	Paticipations []ParticipationResponse
}
type TurnamentResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date" `
	Place     string     `json:"place" `
	Address   string     `json:"address" `
	Village   string     `json:"village" `
	Prize     uint       `json:"prize" `
	Slot      uint       `json:"slot" `
	Champion  string     `json:"champion" `
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

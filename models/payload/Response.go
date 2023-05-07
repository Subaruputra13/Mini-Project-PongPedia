package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProfileResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Player   PlayerResponse
}

type PlayerResponse struct {
	ID        uint   `json:"-"`
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
	ID        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	Location  string `json:"location" form:"location"`
	Slot      int    `json:"slot" form:"slot"`
}

type ParticipationResponse struct {
	PlayerID int `json:"player_id"`
	Player   PlayerResponse
}

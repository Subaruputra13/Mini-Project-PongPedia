package payload

type LoginRequest struct {
	EmailOrUsername string `json:"email_or_username" form:"email_or_username" validate:"required"`
	Password        string `json:"password" form:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserRequest struct {
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type DeleteUserRequest struct {
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type CreateUpdatePlayerRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Age       uint   `json:"age" form:"age" validate:"required"`
	BirthDate string `json:"birth_date" form:"birth_date" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	UserID    uint   `json:"user_id" form:"user_id"`
}

type TurnamentRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	StartDate string `json:"start_date" form:"start_date" validate:"required"`
	EndDate   string `json:"end_date" form:"end_date" validate:"required"`
	Location  string `json:"location" form:"location" validate:"required"`
	Slot      uint   `json:"slot" form:"slot" validate:"required"`
}

type RegisterTurnamentRequest struct {
	PlayerID    int `json:"player_id" form:"player_id"`
	TurnamentID int `json:"turnament_id" form:"turnament_id" validate:"required"`
}

type CreateMatchRequest struct {
	MatchName      string `json:"match_name" form:"match_name" validate:"required"`
	MatchDate      string `json:"match_date" form:"match_date" validate:"required"`
	Player_1       uint   `json:"player_1" form:"player_1" validate:"required"`
	Player_2       uint   `json:"player_2" form:"player_2" validate:"required"`
	Player_1_Score uint   `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint   `json:"player_2_score" form:"player_2_score"`
	TurnamentID    uint   `json:"turnament_id" form:"turnament_id" validate:"required"`
}

type UpdateMatchRequest struct {
	Player_1_Score uint `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint `json:"player_2_score" form:"player_2_score"`
}

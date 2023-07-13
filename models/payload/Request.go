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
	FirstName string `json:"first_name" form:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" validate:"required"`
	Age       uint   `json:"age" form:"age" validate:"required"`
	BirthDate string `json:"birth_date" form:"birth_date" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	Styles    string `json:"styles" form:"styles" validate:"required"`
	UserID    uint   `json:"user_id" form:"user_id"`
}

type TurnamentRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	StartDate string `json:"start_date" form:"start_date" validate:"required"`
	EndDate   string `json:"end_date" form:"end_date" validate:"required"`
	Place     string `json:"place" form:"place" validate:"required"`
	Address   string `json:"address" form:"address" validate:"required"`
	Village   string `json:"village" form:"village" validate:"required"`
	Prize     uint   `json:"prize" form:"prize" validate:"required"`
	Champion  string `json:"champion" form:"champion"`
}
type UpdateTurnamentRequest struct {
	Name     string `json:"name" form:"name"`
	Place    string `json:"place" form:"place"`
	Address  string `json:"address" form:"address"`
	Village  string `json:"village" form:"village"`
	Prize    uint   `json:"prize" form:"prize"`
	Champion string `json:"champion" form:"champion"`
}

type RegisterTurnamentRequest struct {
	PlayerID    int `json:"player_id" form:"player_id"`
	TurnamentID int `json:"turnament_id" form:"turnament_id" validate:"required"`
}

type CreateMatchRequest struct {
	MatchType      string `json:"match_type" form:"match_type" validate:"required"`
	MatchDate      string `json:"match_date" form:"match_date" validate:"required"`
	Player_1ID     uint   `json:"player_1ID" form:"player_1ID" validate:"required"`
	Player_2ID     uint   `json:"player_2ID" form:"player_2ID" validate:"required"`
	Player_1_Score uint   `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint   `json:"player_2_score" form:"player_2_score"`
	TurnamentID    uint   `json:"turnament_id" form:"turnament_id" validate:"required"`
}

type UpdateMatchRequest struct {
	Player_1_Score uint `json:"player_1_score" form:"player_1_score"`
	Player_2_Score uint `json:"player_2_score" form:"player_2_score"`
}

package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProfileResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Player   PlayerResponse
}

type PlayerResponse struct {
	ID        int    `json:"-"`
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

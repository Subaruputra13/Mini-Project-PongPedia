package payload

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

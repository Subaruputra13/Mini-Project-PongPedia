package payload

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type CreateUpdatePlayerRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Age       int    `json:"age" form:"age" validate:"required"`
	BirthDate string `json:"birth_date" form:"birth_date" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	UserID    int    `json:"user_id" form:"user_id"`
}

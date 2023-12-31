package user

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Username string `validate:"required,max=200,min=2" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,max=200,min=2" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UsersResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

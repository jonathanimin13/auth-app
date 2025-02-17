package dto

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	AccesToken string `json:"access_token"`
}

type UserResponseBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type RegisterRequestBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}
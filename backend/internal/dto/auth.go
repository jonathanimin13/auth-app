package dto

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	Username   string `json:"username"`
	AccesToken string `json:"access_token"`
}
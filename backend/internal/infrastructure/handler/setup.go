package handler

import (
	"auth-app/internal/auth/handler"
	"auth-app/internal/auth/repo"
	"auth-app/internal/auth/usecase"
	"auth-app/pkg/bcrypt"
	"auth-app/pkg/jwt"
	"database/sql"
)

type Handler struct {
	AuthHandler handler.AuthHandler
}


func NewHandler(db *sql.DB) *Handler {
	bcrypt := bcrypt.NewBcrypt()
	jwt := jwt.NewJWT()

	authRepo := repo.NewAuthRepo(db)

	authUsecase := usecase.NewAuthUsecase(authRepo, bcrypt, jwt)

	authHandler := handler.NewAuthHandler(authUsecase)

	return &Handler{
		AuthHandler: authHandler,
	}
}
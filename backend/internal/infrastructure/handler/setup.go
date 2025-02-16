package handler

import (
	"auth-app/internal/auth/handler"
	"auth-app/internal/auth/repo"
	"auth-app/internal/auth/usecase"
	"auth-app/internal/infrastructure/db"
	"auth-app/pkg/bcrypt"
	"auth-app/pkg/jwt"
	"log"
)

type Handler struct {
	AuthHandler handler.AuthHandler
}


func NewHandler() *Handler {
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()
	
	bcrypt := bcrypt.NewBcrypt()
	jwt := jwt.NewJWT()

	authRepo := repo.NewAuthRepo(db)

	authUsecase := usecase.NewAuthUsecase(authRepo, bcrypt, jwt)

	authHandler := handler.NewAuthHandler(authUsecase)

	return &Handler{
		AuthHandler: authHandler,
	}
}
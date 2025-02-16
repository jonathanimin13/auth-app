package handler

import "auth-app/internal/auth/usecase"

type AuthHandler interface{}

type authHandlerImpl struct {
	u usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) AuthHandler {
	return &authHandlerImpl{
		u: u,
	}
}
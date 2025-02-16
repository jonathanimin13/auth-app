package usecase

import (
	"auth-app/internal/auth/repo"
	"auth-app/pkg/bcrypt"
	"auth-app/pkg/jwt"
)

type AuthUsecase interface{}

type authUsecaseImpl struct {
	r repo.AuthRepo
	b bcrypt.Bcrypt
	j jwt.JWT
}

func NewAuthUsecase(r repo.AuthRepo, b bcrypt.Bcrypt, j jwt.JWT) AuthUsecase {
	return &authUsecaseImpl{
		r: r,
		b: b,
		j: j,
	}
}
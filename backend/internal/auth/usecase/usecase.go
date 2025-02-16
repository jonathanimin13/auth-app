package usecase

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/auth/autherrors"
	"auth-app/internal/auth/repo"
	"auth-app/internal/entity"
	"auth-app/pkg/bcrypt"
	"auth-app/pkg/customerror"
	"auth-app/pkg/jwt"
	"context"
)

type AuthUsecase interface{
	Login(ctx context.Context, loginData *entity.User) (*entity.User, error)
}

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

func (u *authUsecaseImpl) Login(ctx context.Context, loginData *entity.User) (*entity.User, error) {
	isExists, err := u.r.IsEmailExists(ctx, loginData.Email)
	if err != nil {
		return nil, err
	}
	if !isExists {
		return nil, customerror.NewBadRequestError(autherrors.FieldLogin, autherrors.ErrInvalidEmailOrPassword, autherrors.ErrInvalidEmailOrPassword)
	}

	user, err := u.r.FindUserByEmail(ctx, loginData.Email)
	if err != nil {
		return nil, err
	}

	err = u.b.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return nil, customerror.NewBadRequestError(autherrors.FieldLogin, autherrors.ErrInvalidEmailOrPassword, err)
	}

	accessToken, err := u.j.GenerateAccesToken(user.ID)
	if err != nil {
		return nil, customerror.NewInternalServerError(apperrors.FieldServer, apperrors.ErrInternalServer, err)
	}

	user.AccesToken = accessToken

	return user, nil
}
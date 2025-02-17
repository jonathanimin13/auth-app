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
	Login(ctx context.Context, loginData *entity.User) (string, error)
	User(ctx context.Context, userID int) (*entity.User, error)
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

func (u *authUsecaseImpl) Login(ctx context.Context, loginData *entity.User) (string, error) {
	isExists, err := u.r.IsEmailExists(ctx, loginData.Email)
	if err != nil {
		return "", err
	}
	if !isExists {
		return "", customerror.NewBadRequestError(autherrors.FieldLogin, autherrors.ErrInvalidEmailOrPassword, autherrors.ErrInvalidEmailOrPassword)
	}

	user, err := u.r.FindUserByEmail(ctx, loginData.Email)
	if err != nil {
		return "", err
	}

	err = u.b.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return "", customerror.NewBadRequestError(autherrors.FieldLogin, autherrors.ErrInvalidEmailOrPassword, err)
	}

	accessToken, err := u.j.GenerateAccesToken(user.ID)
	if err != nil {
		return "", customerror.NewInternalServerError(apperrors.FieldServer, apperrors.ErrInternalServer, err)
	}

	return accessToken, nil
}

func (u *authUsecaseImpl) User(ctx context.Context, userID int) (*entity.User, error) {
	user, err := u.r.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
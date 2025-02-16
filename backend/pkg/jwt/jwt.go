package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateAccesToken(userID int) (string, error)
}

type jwtImpl struct{}

func NewJWT() JWT {
	return &jwtImpl{}
}

func (j *jwtImpl) GenerateAccesToken(userID int) (string, error) {
	now := time.Now()
	registeredClaims := jwt.RegisteredClaims{
		Issuer:  "e_commerce_app",
		Subject: fmt.Sprint(userID),
		IssuedAt: &jwt.NumericDate{
			Time: now,
		},
		ExpiresAt: &jwt.NumericDate{
			Time: now.Add(24 * time.Hour),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
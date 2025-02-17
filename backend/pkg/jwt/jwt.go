package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateAccesToken(userID int) (string, error)
	ParseAccessToken(tokenStr string) (*jwt.MapClaims, error)
}

type jwtImpl struct{}

func NewJWT() JWT {
	return &jwtImpl{}
}

func (j *jwtImpl) GenerateAccesToken(userID int) (string, error) {
	now := time.Now()
	registeredClaims := jwt.RegisteredClaims{
		Issuer:  os.Getenv("TOKEN_ISSUER"),
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

func (j *jwtImpl) ParseAccessToken(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenStr,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
		jwt.WithIssuer(os.Getenv("TOKEN_ISSUER")),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &claims, nil
}
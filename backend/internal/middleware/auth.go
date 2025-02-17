package middleware

import (
	"auth-app/internal/apperrors"
	"auth-app/pkg/customerror"
	"auth-app/pkg/jwt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {

	log.Println(ctx.Request.Cookies())

	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, apperrors.ErrInvalidToken))
		ctx.Abort()
		return
	}

	authToken := strings.Split(authHeader, " ")
	token := authToken[1]

	var jwt = jwt.NewJWT()

	claims, err := jwt.ParseAccessToken(token)
	if err != nil {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, err))
		ctx.Abort()
		return
	}

	sub, err := claims.GetSubject()
	if err != nil {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, err))
		ctx.Abort()
		return
	}

	subInt, err := strconv.Atoi(sub)
	if err != nil {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, err))
		ctx.Abort()
		return
	}

	if (subInt < 1) {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, apperrors.ErrInvalidToken))
		ctx.Abort()
		return
	}

	ctx.Set("sub", subInt)
	ctx.Next()
}
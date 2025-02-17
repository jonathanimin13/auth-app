package handler

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/auth/converter"
	"auth-app/internal/auth/usecase"
	"auth-app/internal/dto"
	"auth-app/pkg/customerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface{
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
	User(ctx *gin.Context)
}

type authHandlerImpl struct {
	u usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) AuthHandler {
	return &authHandlerImpl{
		u: u,
	}
}

func (h *authHandlerImpl) Login(ctx *gin.Context)  {
	var loginReqBody dto.LoginRequestBody
	err := ctx.ShouldBindBodyWithJSON(&loginReqBody)
	if err != nil {
		ctx.Error(err)
		return
	}

	loginData := converter.LoginConverter{}.DTOToEntity(&loginReqBody)
	c := ctx.Request.Context()

	accessToken, err := h.u.Login(c, loginData)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie("access-token", accessToken, 86400, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "login succes",
	})
}

func (h *authHandlerImpl) Logout(ctx *gin.Context)  {
	ctx.SetCookie("access-token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "logout succes",
	})
}

func (h *authHandlerImpl) VerifyToken(ctx *gin.Context) {
	_, ok := ctx.Get("sub")
	if !ok {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, apperrors.ErrInvalidToken))
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "token valid",
	})
}

func (h *authHandlerImpl) User(ctx *gin.Context) {
	userID, ok := ctx.Get("sub")
	if !ok {
		ctx.Error(customerror.NewUnauthorizedError(apperrors.FieldToken, apperrors.ErrInvalidToken, apperrors.ErrInvalidToken))
		return
	}

	c := ctx.Request.Context()
	user, err := h.u.User(c, userID.(int))
	if err != nil {
		ctx.Error(err)
		return
	}

	userDTO := converter.UserConverter{}.EntityToDTO(user)

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "token valid",
		Data: userDTO,
	})
}
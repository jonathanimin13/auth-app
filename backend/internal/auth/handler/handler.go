package handler

import (
	"auth-app/internal/auth/converter"
	"auth-app/internal/auth/usecase"
	"auth-app/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface{
	Login(ctx *gin.Context)
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

	user, err := h.u.Login(c, loginData)
	if err != nil {
		ctx.Error(err)
		return
	}

	loginResBody := converter.LoginConverter{}.EntityToDTO(*user)

	ctx.JSON(http.StatusOK, dto.Response{
		Message: "login succes",
		Data: loginResBody,
	})
}
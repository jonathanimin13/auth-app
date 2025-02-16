package router

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/infrastructure/handler"
	"auth-app/internal/middleware"
	"auth-app/pkg/customerror"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger)
	r.Use(middleware.Error)

	setupNoRoute(r)

	baseEndpoint := r.Group("/api")
	setupAuthRoute(baseEndpoint, handler)

	return r
}

func setupAuthRoute(baseEndpoint *gin.RouterGroup, handler *handler.Handler) {
	authGroup := baseEndpoint.Group("/auth")
	{
		authGroup.POST("/login", handler.AuthHandler.Login)
	}
}

func setupNoRoute(r *gin.Engine) {
	r.NoRoute(
		func(ctx *gin.Context) {
			ctx.Error(customerror.NewNotFoundError(apperrors.FieldNotFound, apperrors.ErrEndpointNotFound, apperrors.ErrEndpointNotFound))
		},
	)
}
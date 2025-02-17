package router

import (
	"auth-app/internal/apperrors"
	"auth-app/internal/infrastructure/handler"
	"auth-app/internal/middleware"
	"auth-app/pkg/customerror"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger)
	r.Use(middleware.Error)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow only specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow credentials (cookies, authorization headers)
		MaxAge:           12 * time.Hour, // Cache preflight requests for 12 hours
}))

	setupNoRoute(r)

	baseEndpoint := r.Group("/api")
	setupAuthRoute(baseEndpoint, handler)

	return r
}

func setupAuthRoute(baseEndpoint *gin.RouterGroup, handler *handler.Handler) {
	authGroup := baseEndpoint.Group("/auth")
	{
		authGroup.POST("/login", handler.AuthHandler.Login)
		authGroup.POST("/logout", handler.AuthHandler.Logout)
		authGroup.GET("/verify-token", middleware.Auth, handler.AuthHandler.VerifyToken)
		authGroup.GET("/user", middleware.Auth, handler.AuthHandler.User)
	}
}

func setupNoRoute(r *gin.Engine) {
	r.NoRoute(
		func(ctx *gin.Context) {
			ctx.Error(customerror.NewNotFoundError(apperrors.FieldNotFound, apperrors.ErrEndpointNotFound, apperrors.ErrEndpointNotFound))
		},
	)
}
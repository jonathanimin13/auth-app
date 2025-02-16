package router

import (
	"auth-app/internal/infrastructure/handler"
	"auth-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger)
	r.Use(middleware.Error)

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
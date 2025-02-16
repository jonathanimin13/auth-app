package router

import (
	"auth-app/internal/auth/middleware"
	"auth-app/internal/infrastructure/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger)

	return r
}
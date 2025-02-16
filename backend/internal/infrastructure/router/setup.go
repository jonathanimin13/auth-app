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

	return r
}
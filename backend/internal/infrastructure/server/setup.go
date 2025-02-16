package server

import (
	"auth-app/internal/config"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer(router *gin.Engine) *http.Server {
	serverPort := os.Getenv("SERVER_PORT")

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", serverPort),
		Handler:      router,
		ReadTimeout:  config.Timeout * time.Minute,
		WriteTimeout: config.Timeout * time.Minute,
	}
}
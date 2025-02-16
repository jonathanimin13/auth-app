package main

import (
	"auth-app/internal/infrastructure/handler"
	"auth-app/internal/infrastructure/router"
	"auth-app/pkg/logger"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env: %v", err)
	}

	logger.SetLogger(logger.NewLogrusLogger())

	handler := handler.NewHandler()

	router := router.NewRouter(handler)
}
package main

import (
	"auth-app/internal/infrastructure/db"
	"auth-app/pkg/logger"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("env: %v", err)
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()

	logger.SetLogger(logger.NewLogrusLogger())
}
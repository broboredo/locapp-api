package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	env := "." + os.Getenv("GO_ENV")

	err := godotenv.Load(".env" + env)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

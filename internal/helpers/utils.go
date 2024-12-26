package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using default values where necessary.")
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

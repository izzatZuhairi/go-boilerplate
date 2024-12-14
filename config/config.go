package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	return os.Getenv(key)
}

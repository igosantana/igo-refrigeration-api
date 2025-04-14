package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

var AppConfig Config

func LoadEnv() {
	AppConfig = Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}

	if AppConfig.DatabaseURL == "" {
		log.Fatal("DATABASE_URL não definida no ambiente")
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

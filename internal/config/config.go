package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type AppConfig struct {
	Env  string
	Port string
}

type Config struct {
	Database DatabaseConfig
	App      AppConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("config: .env file not loaded, using environment variables: %v", err)
	}

	return &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "PRODUCTION"),
			Port: getEnv("PORT", "3000"),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  getEnv("DB_SSL_MODE", "require"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
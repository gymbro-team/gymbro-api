package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User     string
	Password string
	Name     string
}

type Config struct {
	Database      DatabaseConfig
	ServerAddress string
	Token         string
}

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return value
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConfig := DatabaseConfig{
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		Name:     getEnv("DB_NAME", ""),
	}

	return &Config{
		Database:      dbConfig,
		ServerAddress: getEnv("SERVER_ADDRESS", ":3000"),
		Token:         getEnv("TOKEN", ""),
	}
}

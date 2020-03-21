package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPort     string
	DBUser     string
	DBPass     string
	DBHost     string
	DBDatabase string

	ServerHost string
	ServerPort string
}

func StartConfigs() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	return &Config{
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBPass:     os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

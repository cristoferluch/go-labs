package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Token string
}

func ParseConfig() (*Config, error) {

	if err := godotenv.Load(".env"); err != nil {
		return nil, errors.New("not found .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("BOT_TOKEN is empty")
	}

	return &Config{Token: token}, nil
}

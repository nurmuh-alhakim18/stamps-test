package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("failed to load .env")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("WEATHER_API_KEY must be set")
	}

	return &Config{
		ApiKey: apiKey,
	}, nil
}

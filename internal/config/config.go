package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WeatherAPIKey string
	HTTPPort      string
}

// LoadConfig loads environment variables from a .env file or the system environment
func LoadConfig() (*Config, error) {
	// Attempt to load a .env file, but don't fail if it doesn't exist
	// (for example, in production/Cloud Run where environment variables are set via the console)
	_ = godotenv.Load()

	cfg := &Config{
		WeatherAPIKey: os.Getenv("WEATHER_API_KEY"),
		HTTPPort:      os.Getenv("PORT"),
	}

	if cfg.WeatherAPIKey == "" {
		return nil, errors.New("WEATHER_API_KEY is missing")
	}

	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8080"
	}

	return cfg, nil
}
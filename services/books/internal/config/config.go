package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	Addr   string
}

func New() *Config {
	godotenv.Load()
	return &Config{
		APIKey: os.Getenv("API_KEY"),
		Addr:   os.Getenv("ADDR"),
	}
}

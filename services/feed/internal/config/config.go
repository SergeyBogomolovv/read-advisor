package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AmqpURL      string
	PostgresURL  string
	BooksGrpcURL string
	Addr         string
}

func New() Config {
	godotenv.Load()
	return Config{
		AmqpURL:      os.Getenv("AMQP_URL"),
		PostgresURL:  os.Getenv("POSTGRES_URL"),
		BooksGrpcURL: os.Getenv("BOOKS_GRPC_URL"),
		Addr:         os.Getenv("ADDR"),
	}
}

package configs

import (
	"github.com/NRKA/home-service/pkg/postgres"
	"github.com/joho/godotenv"
	"os"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

func FromEnv() (postgres.DatabaseConfig, error) {
	if err := godotenv.Load(); err != nil {
		return postgres.DatabaseConfig{}, err
	}

	return postgres.DatabaseConfig{
		Host:     os.Getenv(dbHost),
		Port:     os.Getenv(dbPort),
		User:     os.Getenv(dbUser),
		Password: os.Getenv(dbPassword),
		Name:     os.Getenv(dbName),
	}, nil
}

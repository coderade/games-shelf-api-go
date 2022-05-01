package config

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	Env      string
	Db       DatabaseConfig
	Secret   string
	LogLevel string
}

type DatabaseConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	port := getEnv("PORT", "4000")
	env := getEnv("ENV", "development")
	dbDsn := getEnv("DB_DATA_SOURCE", "postgres://admin@localhost/games_shelf?sslmode=disable")
	secret := getEnv("APP_SECRET", "games-shelf-api-secret")
	logLevel := getEnv("LOG_LEVEL", "info")

	return &Config{
		Port:     port,
		Env:      env,
		Db:       DatabaseConfig{Dsn: dbDsn},
		Secret:   secret,
		LogLevel: logLevel,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Environment variable %s not set, using default: %s", key, fallback)
	return fallback
}

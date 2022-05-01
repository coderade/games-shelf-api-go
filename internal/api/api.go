package api

import (
	"context"
	"database/sql"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/models"
	"log"
	"os"
	"time"
)

type Server struct {
	Config config.Config
	Shelf  *models.Shelf
}

func (api *Server) Initialize() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	api.Config = getConfigVariables()
	db, err := openDBConnection(api.Config)
	if err != nil {
		logger.Fatal(err)
	}

	api.Shelf = models.NewShelf(db)

}

func getConfigVariables() config.Config {
	var cfg config.Config

	cfg.Port = getEnv("PORT", "4000")      // Server port to list on
	cfg.Env = getEnv("ENV", "development") // Application environment (development|production)
	cfg.Db.Dsn = getEnv("DB_DATA_SOURCE",
		"postgres://admin@localhost/games_shelf?sslmode=disable") // Postgres Data Source
	cfg.Secret = getEnv("APP_SECRET", "games-shelf-api-secret") // Application secret
	config.SetConfig(cfg)
	return cfg
}

func openDBConnection(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil

}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

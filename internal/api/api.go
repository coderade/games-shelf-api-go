package api

import (
	"context"
	"database/sql"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/logger"
	"games-shelf-api-go/internal/models"
	"time"
)

type Server struct {
	Config *config.Config
	Shelf  *models.Shelf
	Logger *logger.Logger
}

func (s *Server) Initialize(cfg *config.Config, log *logger.Logger) {
	s.Config = cfg
	s.Logger = log

	// Initialize DB and other services here
	db, err := openDBConnection(cfg, log)
	if err != nil {
		log.Fatal(err)
	}
	s.Shelf = models.NewShelf(db)
}

func openDBConnection(cfg *config.Config, log *logger.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

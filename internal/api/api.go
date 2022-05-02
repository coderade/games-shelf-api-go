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
	DB     *sql.DB
}

// Initialize initializes the server with the provided configuration and logger.
func (s *Server) Initialize(cfg *config.Config, log *logger.Logger) {
	s.Config = cfg
	s.Logger = log

	db, err := openDBConnection(cfg, log)
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}
	s.DB = db
	s.Shelf = models.NewShelf(db)
}

// openDBConnection opens a database connection and pings it to ensure it's reachable.
func openDBConnection(cfg *config.Config, log *logger.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		log.Error("Error opening database connection: ", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Error("Error pinging database: ", err)
		return nil, err
	}

	log.Info("Successfully connected to the database")
	return db, nil
}

// Close closes the database connection.
func (s *Server) Close() {
	if s.DB != nil {
		err := s.DB.Close()
		if err != nil {
			s.Logger.Error("Error closing database connection: ", err)
		} else {
			s.Logger.Info("Database connection closed")
		}
	}
}

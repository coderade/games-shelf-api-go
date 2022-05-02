package api

import (
	"context"
	"database/sql"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/logger"
	"games-shelf-api-go/internal/repository"
	"time"
)

type Server struct {
	Config   *config.Config
	Shelf    *repository.Shelf
	Logger   *logger.Logger
	DBHelper *db.DBHelper
}

// Initialize initializes the server with the provided configuration and logger.
func (s *Server) Initialize(cfg *config.Config, log *logger.Logger) {
	s.Config = cfg
	s.Logger = log

	dbConn, err := openDBConnection(cfg, log)
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}

	s.DBHelper = db.NewDBHelper(dbConn)
	s.Shelf = repository.NewShelf(dbConn)
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
	if s.DBHelper != nil {
		err := s.DBHelper.Close()
		if err != nil {
			s.Logger.Error("Error closing database connection: ", err)
		} else {
			s.Logger.Info("Database connection closed")
		}
	}
}

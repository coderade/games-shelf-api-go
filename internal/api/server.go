package api

import (
	"context"
	"database/sql"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/logger"
	"games-shelf-api-go/internal/repository"
	"net/http"
	"time"
)

type Server struct {
	Config   *config.Config
	Shelf    *repository.Shelf
	Logger   *logger.Logger
	DBHelper db.Database
	Router   *http.ServeMux
}

// Initialize initializes the server with the provided configuration, logger, and database.
// If dbHelper is provided, it uses that; otherwise, it opens a new database connection.
func (s *Server) Initialize(cfg *config.Config, log *logger.Logger, dbHelper db.Database) {
	s.Config = cfg
	s.Logger = log

	if dbHelper != nil {
		s.DBHelper = dbHelper
	} else {
		var err error
		s.DBHelper, err = s.openDBConnection(cfg, log)
		if err != nil {
			log.Fatal("Failed to open database connection: ", err)
		}
	}

	s.Shelf = repository.NewShelf(s.DBHelper)
}

// openDBConnection opens a database connection and pings it to ensure it's reachable.
func (s *Server) openDBConnection(cfg *config.Config, log *logger.Logger) (db.Database, error) {
	sqlDB, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		log.Error("Error opening database connection: ", err)
		return nil, err
	}

	database := &db.SQLDatabase{DB: sqlDB}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = database.PingContext(ctx)
	if err != nil {
		log.Error("Error pinging database: ", err)
		return nil, err
	}

	log.Info("Successfully connected to the database")
	return database, nil
}

// Close closes the database connection.
func (s *Server) Close() {
	if s.DBHelper != nil {
		s.Logger.Info("Closing database connection...")
		err := s.DBHelper.Close()
		if err != nil {
			s.Logger.Error("Error closing database connection: ", err)
		} else {
			s.Logger.Info("Database connection closed")
		}
	}
}

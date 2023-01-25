package main

import (
	"context"
	"database/sql"
	"fmt"
	"games-shelf-api-go/internal/api"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

var log *logger.Logger

func startServer() *http.Server {
	// Load configuration
	cfg := config.LoadConfig()
	log = logger.NewLogger(cfg.LogLevel)

	// Initialize server
	app := api.Server{}
	dbConn, err := openDBConnection(cfg, log)
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}
	app.Initialize(cfg, log, dbConn)

	port := app.Config.Port
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Infof("Starting server on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v\n", port, err)
		}
	}()

	return server
}

// openDBConnection opens a database connection and pings it to ensure it's reachable.
func openDBConnection(cfg *config.Config, log *logger.Logger) (db.Database, error) {
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

func main() {
	server := startServer()

	// Channel to listen for interrupt or terminate signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-done
	log.Info("Shutting down server...")

	// Create a deadline to wait for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info("Server exiting")
}

package main

import (
	"context"
	"fmt"
	"games-shelf-api-go/internal/api"
	"games-shelf-api-go/internal/config"
	"games-shelf-api-go/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log := logger.NewLogger(cfg.LogLevel)

	// Initialize server
	app := api.Server{}
	app.Initialize(cfg, log)

	port := app.Config.Port
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Channel to listen for interrupt or terminate signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Infof("Starting server on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v\n", port, err)
		}
	}()

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

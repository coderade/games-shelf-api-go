package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"games-shelf-api-go/cmd/models"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port   int
	env    string
	secret string
	db     struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	shelf  models.Shelf
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to list on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://admin@localhost/games_shelf?sslmode=disable", "Postgress Data Source")
	flag.Parse()

	cfg.secret = os.Getenv("APP_SECRET")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, error := openDBConnection(cfg)
	if error != nil {
		logger.Fatal(error)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		shelf:  models.NewShelf(db),
	}

	fmt.Println("Running")

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err := server.ListenAndServe()

	if err != nil {
		logger.Println(err)
	}

}

func openDBConnection(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
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

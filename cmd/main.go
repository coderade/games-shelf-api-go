package main

import (
	"flag"
	"fmt"
	"games-shelf-api-go/cmd/api"
	"games-shelf-api-go/cmd/config"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to list on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "postgres://admin@localhost/games_shelf?sslmode=disable", "Postgres Data Source")
	flag.Parse()

	cfg.Secret = os.Getenv("APP_SECRET")
	fmt.Println("Running")

	app := api.Server{}

	app.Initialize(cfg)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	println("Starting server on port", cfg.Port)

	err := server.ListenAndServe()

	if err != nil {
		println(err)
	}

}

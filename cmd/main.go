package main

import (
	"fmt"
	"games-shelf-api-go/cmd/api"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	app := api.Server{}
	app.Initialize()

	port := app.Config.Port
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	println("Starting server on port", port)

	err := server.ListenAndServe()

	if err != nil {
		println(err)
	}

}

package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.getGame)
	router.HandlerFunc(http.MethodGet, "/v1/games", app.getAllGames)

	return app.enableCORS(router)
}

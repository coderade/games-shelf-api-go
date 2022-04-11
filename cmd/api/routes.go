package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", app.SignIn)

	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.getGame)
	router.HandlerFunc(http.MethodPut, "/v1/games/edit/:id", app.editGame)
	router.HandlerFunc(http.MethodDelete, "/v1/games/delete/:id", app.deleteGame)
	router.HandlerFunc(http.MethodPost, "/v1/games/add", app.addGame)
	router.HandlerFunc(http.MethodGet, "/v1/games", app.getAllGames)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/platforms", app.GetAllPlatforms)

	return app.enableCORS(router)
}

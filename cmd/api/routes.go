package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := context.WithValue(request.Context(), "params", params)
		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.validateJWTToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// public routes
	router.HandlerFunc(http.MethodGet, "/v1/games", app.getAllGames)
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", app.getGame)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/platforms", app.GetAllPlatforms)

	//auth routes
	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", app.SignIn)

	// private routes
	router.PUT("/v1/games/edit/:id", app.wrap(secure.ThenFunc(app.editGame)))
	router.DELETE("/v1/games/delete/:id", app.wrap(secure.ThenFunc(app.deleteGame)))
	router.POST("/v1/games/add", app.wrap(secure.ThenFunc(app.addGame)))

	return app.enableCORS(router)
}

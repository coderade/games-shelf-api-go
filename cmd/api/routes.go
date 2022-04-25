package api

import (
	"context"
	"games-shelf-api-go/cmd/api/handlers"
	"games-shelf-api-go/cmd/models"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (api *Server) wrap(next http.Handler) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := context.WithValue(request.Context(), "params", params)
		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}

func (api *Server) Routes() http.Handler {

	router := httprouter.New()
	secure := alice.New(api.validateJWTToken)

	router.HandlerFunc(http.MethodGet, "/status", handlers.StatusHandler)

	// public routes
	router.HandlerFunc(http.MethodGet, "/v1/games", api.handleRequest(handlers.GetAllGames))
	router.HandlerFunc(http.MethodGet, "/v1/games/:id", api.handleRequest(handlers.GetGame))
	router.HandlerFunc(http.MethodGet, "/v1/genres", api.handleRequest(handlers.GetAllGenres))
	router.HandlerFunc(http.MethodGet, "/v1/platforms", api.handleRequest(handlers.GetAllPlatforms))

	//auth routes
	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", handlers.SignIn)

	// private routes
	router.PUT("/v1/games/edit/:id", api.wrap(secure.ThenFunc(api.handleRequest(handlers.EditGame))))
	router.DELETE("/v1/games/delete/:id", api.wrap(secure.ThenFunc(api.handleRequest(handlers.DeleteGame))))
	router.POST("/v1/games/add", api.wrap(secure.ThenFunc(api.handleRequest(handlers.AddGame))))

	// graphql routes
	router.HandlerFunc(http.MethodPost, "/v1/graphql", api.handleRequest(handlers.GamesGraphQL))

	return api.enableCORS(router)

}

type RequestHandlerFunction func(shelf *models.Shelf, w http.ResponseWriter, r *http.Request)

func (api *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(api.Shelf, w, r)
	}
}

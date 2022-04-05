package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getGame(writer http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(writer, err)
		return
	}

	game, err := app.shelf.GetGameById(id)

	err = app.writeJSON(writer, http.StatusOK, game, "game")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

}

func (app *application) getAllGames(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	genreId, err := strconv.Atoi(queryValues.Get("genre_id"))
	platformId, err := strconv.Atoi(queryValues.Get("platform_id"))
	games, err := app.shelf.GetAllGames(genreId, platformId)
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	err = app.writeJSON(writer, http.StatusOK, games, "games")
	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}

//func (app *application) getAllGamesByGenre(writer http.ResponseWriter, request *http.Request) {
//	queryValues := request.URL.Query()
//	genreId, err := strconv.Atoi(queryValues.Get("genre_id"))
//	games, err := app.shelf.GetAllGames(genreId, 0)
//	if err != nil {
//		app.errorJSON(writer, err)
//		return
//	}
//
//	err = app.writeJSON(writer, http.StatusOK, games, "games")
//	if err != nil {
//		app.errorJSON(writer, err)
//		return
//	}
//}
//
//func (app *application) getAllGamesByPlatform(writer http.ResponseWriter, request *http.Request) {
//	queryValues := request.URL.Query()
//	platformId, err := strconv.Atoi(queryValues.Get("platform_id"))
//	games, err := app.shelf.GetAllGames(0, platformId)
//	if err != nil {
//		app.errorJSON(writer, err)
//		return
//	}
//
//	err = app.writeJSON(writer, http.StatusOK, games, "games")
//	if err != nil {
//		app.errorJSON(writer, err)
//		return
//	}
//}

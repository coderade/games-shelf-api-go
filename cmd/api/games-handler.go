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

	app.logger.Println("ID is:", id)

	game, err := app.shelf.GetGameById(id)

	err = app.writeJSON(writer, http.StatusOK, game, "game")

	if err != nil {
		app.errorJSON(writer, err)
		return
	}

}

func (app *application) getAllGames(writer http.ResponseWriter, reader *http.Request) {
	games, err := app.shelf.GetAllGames()
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

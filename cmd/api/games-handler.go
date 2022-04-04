package main

import (
	"encoding/json"
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

}

func (app *application) getAllGames(writer http.ResponseWriter, reader *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	j, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(j)
	if err != nil {
		app.logger.Println(err)
	}
}

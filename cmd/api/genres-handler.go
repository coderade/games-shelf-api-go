package main

import (
	"net/http"
)

func (app *application) getAllGenres(writer http.ResponseWriter, request *http.Request) {
	genres, err := app.shelf.GetAllGenres()
	if err != nil {
		app.errorJSON(writer, err, http.StatusBadRequest)
		return
	}

	err = app.writeJSON(writer, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(writer, err, http.StatusBadRequest)
		return
	}
}

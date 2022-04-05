package main

import (
	"net/http"
)

func (app *application) getAllGenres(writer http.ResponseWriter, reader *http.Request) {
	genres, err := app.shelf.GetAllGenres()
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	err = app.writeJSON(writer, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}

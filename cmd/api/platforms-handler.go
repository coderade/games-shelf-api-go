package main

import (
	"net/http"
)

func (app *application) GetAllPlatforms(writer http.ResponseWriter, reader *http.Request) {
	platforms, err := app.shelf.GetAllPlatforms()
	if err != nil {
		app.errorJSON(writer, err)
		return
	}

	err = app.writeJSON(writer, http.StatusOK, platforms, "platforms")
	if err != nil {
		app.errorJSON(writer, err)
		return
	}
}

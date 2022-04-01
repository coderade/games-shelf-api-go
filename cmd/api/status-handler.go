package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(writer http.ResponseWriter, reader *http.Request) {
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

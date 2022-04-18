package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(writer http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	_, err = writer.Write(js)
	return nil
}
func (app *application) errorJSON(writer http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	e := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(writer, statusCode, e, "error")
}

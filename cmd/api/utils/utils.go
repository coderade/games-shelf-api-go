package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(writer http.ResponseWriter, status int, data interface{}, wrap string) error {
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
func WriteErrorJson(writer http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	errorMessage := err.Error()

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	e := jsonError{
		Message: errorMessage,
	}

	println(errorMessage)
	WriteJson(writer, statusCode, e, "error")
}

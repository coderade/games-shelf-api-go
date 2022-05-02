package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON writes a JSON response to the writer.
func WriteJSON(writer http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := map[string]interface{}{wrap: data}

	js, err := json.Marshal(wrapper)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	_, err = writer.Write(js)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
	return err
}

// WriteErrorJSON writes a JSON error response to the writer.
func WriteErrorJSON(writer http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	errorMessage := err.Error()

	type jsonError struct {
		Message string `json:"message"`
	}

	e := jsonError{
		Message: errorMessage,
	}

	log.Printf("Error: %s", errorMessage)
	WriteJSON(writer, statusCode, e, "error")
}

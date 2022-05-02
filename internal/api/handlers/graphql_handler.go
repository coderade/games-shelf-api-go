package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"games-shelf-api-go/internal/repository"
	"games-shelf-api-go/internal/utils"
	"io"
	"log"
	"net/http"

	graphqlschema "games-shelf-api-go/internal/service/graphql"

	"github.com/graphql-go/graphql"
)

// GamesGraphQL handles GraphQL queries for games
func GamesGraphQL(shelf *repository.Shelf, writer http.ResponseWriter, request *http.Request) {
	schema, err := graphqlschema.NewSchema(shelf)
	if err != nil {
		log.Printf("Error creating GraphQL schema: %v", err)
		utils.WriteErrorJSON(writer, errors.New("failed to create the GraphQL schema"), http.StatusInternalServerError)
		return
	}

	query, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		utils.WriteErrorJSON(writer, errors.New("invalid request body"), http.StatusBadRequest)
		return
	}

	result := executeQuery(string(query), schema)
	if len(result.Errors) > 0 {
		log.Printf("GraphQL errors: %v", result.Errors)
		utils.WriteErrorJSON(writer, errors.New(fmt.Sprintf("GraphQL errors: %+v", result.Errors)), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(result.Data); err != nil {
		log.Printf("Error writing response: %v", err)
		utils.WriteErrorJSON(writer, errors.New("failed to write response"), http.StatusInternalServerError)
	}
}

// executeQuery executes the given GraphQL query on the provided schema
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	return graphql.Do(params)
}

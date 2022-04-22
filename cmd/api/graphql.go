package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"games-shelf-api-go/cmd/models"
	"github.com/graphql-go/graphql"
	"io"
	"net/http"
	"strings"
)

var games []models.Game

var graphQLFields = graphql.Fields{
	"game": &graphql.Field{
		Type:        gameType,
		Description: "Get a game by the ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, game := range games {
					if game.ID == id {
						return game, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(gameType),
		Description: "Get all games",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return games, nil
		},
	},
	"search": &graphql.Field{
		Type:        graphql.NewList(gameType),
		Description: "Search movies by title",
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var result []*models.Game
			search, ok := params.Args["titleContains"].(string)
			if ok {
				for _, currentGame := range games {
					if strings.Contains(currentGame.Title, search) {
						result = append(result, &currentGame)
					}
				}
			}
			return result, nil
		},
	},
}

var gameType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Game",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"publisher": &graphql.Field{
				Type: graphql.String,
			},
			"rating": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

func (app *application) gamesGraphQL(writer http.ResponseWriter, request *http.Request) {
	games, _ = app.shelf.GetAllGames(0, 0)

	q, err := io.ReadAll(request.Body)
	if err != nil {
		app.errorJSON(writer, err)
		return
	}
	query := string(q)
	app.logger.Println(query)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphQLFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(writer, errors.New("failed to create the graphQL schema"))
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		app.errorJSON(writer, errors.New(fmt.Sprintf("failed: %+v", resp.Errors)))
	}

	js, _ := json.MarshalIndent(resp.Data, "", " ")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(js)

}

package graphql

import (
	"games-shelf-api-go/internal/models"
	"strings"

	"github.com/graphql-go/graphql"
)

var games []models.Game

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
		Description: "Search games by title",
		Args: graphql.FieldConfigArgument{
			"titleContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var result []models.Game
			var gameFound models.Game
			search, ok := params.Args["titleContains"].(string)
			if ok {
				for _, currentGame := range games {
					if strings.Contains(currentGame.Title, search) {
						gameFound = currentGame
						result = append(result, gameFound)
					}
				}
			}

			return result, nil
		},
	},
}

// NewSchema creates a new GraphQL schema for games
func NewSchema(shelf *models.Shelf) (graphql.Schema, error) {
	var err error
	games, err = shelf.GetAllGames(0, 0)
	if err != nil {
		return graphql.Schema{}, err
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphQLFields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}

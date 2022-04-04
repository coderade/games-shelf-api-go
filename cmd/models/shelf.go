package models

import (
	"context"
	"database/sql"
	"time"
)

type Shelf struct {
	DB *sql.DB
}

func NewShelf(db *sql.DB) Shelf {
	return Shelf{
		DB: db,
	}
}

// GetGameById returns a game specified by the ID and an error, if any
func (shelf *Shelf) GetGameById(id int) (*Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, publisher, rating, created_at, updated_at 
			FROM games WHERE id = $1`

	row := shelf.DB.QueryRowContext(ctx, query, id)
	var game Game

	err := row.Scan(
		&game.ID,
		&game.Title,
		&game.Description,
		&game.Year,
		&game.Publisher,
		&game.Rating,
		&game.CreatedAt,
		&game.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	id = 1
	query = `SELECT gn.genre_name,
				   p.platform_name,
       			   p.generation
			FROM games g
					LEFT JOIN games_genres gg on g.id = gg.game_id
					 LEFT JOIN games_platforms gp on g.id = gp.game_id
					 LEFT JOIN genres gn on (gn.id = gg.genre_id)
					 LEFT JOIN platforms p on (p.id = gp.platform_id)
			WHERE gg.game_id = $1
			  and gp.game_id = $1
`

	rows, _ := shelf.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var gameGenres []GameGenre
	var gamePlatforms []GamePlatform

	for rows.Next() {
		var gg GameGenre
		var gp GamePlatform
		err := rows.Scan(
			&gg.Genre.Name,
			&gp.Platform.Name,
			&gp.Platform.Generation,
		)

		if err != nil {
			return nil, err
		}

		gameGenres = append(gameGenres, gg)
		gamePlatforms = append(gamePlatforms, gp)
	}

	game.GameGenre = gameGenres
	game.GamePlatform = gamePlatforms
	return &game, nil
}

// GetAllGames returns all games and an error, if any
func (shelf *Shelf) GetAllGames(id int) ([]*Game, error) {
	return nil, nil
}

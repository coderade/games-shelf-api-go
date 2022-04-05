package models

import (
	"context"
	"database/sql"
	"golang.org/x/exp/slices"
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

	gameGenres, gamePlatforms, err := shelf.getGenresAndPlatformsByGameId(id)

	game.GameGenre = gameGenres
	game.GamePlatform = gamePlatforms
	return &game, nil
}

// GetAllGames returns all games and an error, if any
func (shelf *Shelf) GetAllGames() ([]Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, publisher, rating, created_at, updated_at 
			FROM games ORDER BY title`

	rows, err := shelf.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var games []Game

	for rows.Next() {
		var game Game

		err := rows.Scan(
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

		gameGenres, gamePlatforms, err := shelf.getGenresAndPlatformsByGameId(game.ID)
		if err != nil {
			return nil, err
		}

		game.GameGenre = gameGenres
		game.GamePlatform = gamePlatforms

		games = append(games, game)
	}

	return games, nil
}

// GetAllGenres returns all genres and an error, if any
func (shelf *Shelf) GetAllGenres() ([]Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at 
			FROM genres ORDER BY genre_name`

	rows, err := shelf.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var genres []Genre

	for rows.Next() {
		var genre Genre

		err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

// GetAllPlatforms returns all platforms and an error, if any
func (shelf *Shelf) GetAllPlatforms() ([]Platform, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, platform_name, created_at, updated_at 
			FROM platforms ORDER BY platform_name`

	rows, err := shelf.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var platforms []Platform

	for rows.Next() {
		var platform Platform

		err := rows.Scan(
			&platform.ID,
			&platform.Name,
			&platform.CreatedAt,
			&platform.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		platforms = append(platforms, platform)
	}

	return platforms, nil
}

func (shelf *Shelf) getGenresAndPlatformsByGameId(id int) ([]Genre, []Platform, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
       			gn.id, 
       			gn.genre_name,
       			p.id,
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

	var gameGenres []Genre
	var gamePlatforms []Platform

	for rows.Next() {
		var g Genre
		var p Platform
		err := rows.Scan(
			&g.ID,
			&g.Name,
			&p.ID,
			&p.Name,
			&p.Generation,
		)

		if err != nil {
			return nil, nil, err
		}

		// add a platform only if not exists
		if !slices.Contains(gamePlatforms, p) {
			gamePlatforms = append(gamePlatforms, p)
		}

		// add a genre only if not exists
		if !slices.Contains(gameGenres, g) {
			gameGenres = append(gameGenres, g)
		}
	}

	return gameGenres, gamePlatforms, nil
}

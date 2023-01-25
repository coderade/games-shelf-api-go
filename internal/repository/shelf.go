package repository

import (
	"context"
	"fmt"
	"games-shelf-api-go/internal/db"
	"games-shelf-api-go/internal/models"
	"time"

	"golang.org/x/exp/slices"
)

// Shelf represents the repository for accessing game data.
type Shelf struct {
	DB db.Database
}

// NewShelf creates a new Shelf repository.
func NewShelf(database db.Database) *Shelf {
	return &Shelf{
		DB: database,
	}
}

// GetGameById returns a game specified by the ID and an error, if any.
func (shelf *Shelf) GetGameById(id int) (*models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, publisher, rawg_id, created_at, updated_at 
			FROM games WHERE id = $1`

	row := shelf.DB.QueryRowContext(ctx, query, id)
	var game models.Game

	err := row.Scan(
		&game.ID,
		&game.Title,
		&game.Description,
		&game.Year,
		&game.Publisher,
		&game.RawgId,
		&game.CreatedAt,
		&game.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	gameGenres, gamePlatforms, err := shelf.getGenresAndPlatformsByGameId(id)
	if err != nil {
		return nil, err
	}

	game.GameGenre = gameGenres
	game.GamePlatform = gamePlatforms
	return &game, nil
}

// GetAllGames returns all games and an error, if any.
func (shelf *Shelf) GetAllGames(genreID int, platformID int) ([]models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	platformWhere := ""
	if platformID != 0 {
		platformWhere = fmt.Sprintf("id IN (SELECT game_id FROM games_platforms WHERE platform_id = %d)", platformID)
	}

	genreWhere := ""
	if genreID != 0 {
		genreWhere = fmt.Sprintf("id IN (SELECT game_id FROM games_genres WHERE genre_id = %d)", genreID)
	}

	query := `SELECT id, title, description, year, publisher, rating, created_at, updated_at 
			FROM games`

	if genreWhere != "" {
		query = query + " WHERE " + genreWhere

		if platformWhere != "" {
			query = query + " AND " + platformWhere
		}
	} else if platformWhere != "" {
		query = query + " WHERE " + platformWhere

		if genreWhere != "" {
			query = query + " AND " + genreWhere
		}
	}

	query = query + " ORDER BY title"

	rows, err := shelf.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game

	for rows.Next() {
		var game models.Game

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

// GetAllGenres returns all genres and an error, if any.
func (shelf *Shelf) GetAllGenres() ([]models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, genre_name, created_at, updated_at 
			FROM genres ORDER BY genre_name`

	rows, err := shelf.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genre

	for rows.Next() {
		var genre models.Genre

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

// GetAllPlatforms returns all platforms and an error, if any.
func (shelf *Shelf) GetAllPlatforms() ([]models.Platform, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, platform_name, created_at, updated_at 
			FROM platforms ORDER BY platform_name`

	rows, err := shelf.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var platforms []models.Platform

	for rows.Next() {
		var platform models.Platform

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

// getGenresAndPlatformsByGameId returns the genres and platforms associated with a game by its ID.
func (shelf *Shelf) getGenresAndPlatformsByGameId(id int) ([]models.Genre, []models.Platform, error) {
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

	var gameGenres []models.Genre
	var gamePlatforms []models.Platform

	for rows.Next() {
		var g models.Genre
		var p models.Platform
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

// AddGame adds a new game to the repository.
func (shelf *Shelf) AddGame(game models.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO public.games ( title, description, year, publisher, rating, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := shelf.DB.ExecContext(ctx, stmt,
		game.Title,
		game.Description,
		game.Year,
		game.Publisher,
		game.Rating,
		game.CreatedAt,
		game.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// EditGame updates an existing game in the repository.
func (shelf *Shelf) EditGame(game models.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `UPDATE public.games SET title = $1, description = $2 , year =$3, publisher = $4, rating = $5,
                      created_at = $6, updated_at = $7 WHERE id = $8`

	_, err := shelf.DB.ExecContext(ctx, stmt,
		game.Title,
		game.Description,
		game.Year,
		game.Publisher,
		game.Rating,
		game.CreatedAt,
		game.UpdatedAt,
		game.ID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteGame deletes a game from the repository by its ID.
func (shelf *Shelf) DeleteGame(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `DELETE FROM public.games WHERE id = $1`

	_, err := shelf.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}

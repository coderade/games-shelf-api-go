package repository

import (
	"games-shelf-api-go/internal/models"
)

type ShelfRepository interface {
	GetGameById(id int) (*models.Game, error)
	GetAllGames(genreID int, platformID int) ([]models.Game, error)
	AddGame(game models.Game) error
	EditGame(game models.Game) error
	DeleteGame(id int) error
	GetGenresAndPlatformsByGameId(id int) ([]models.Genre, []models.Platform, error)
	GetAllGenres() ([]models.Genre, error)
	GetAllPlatforms() ([]models.Platform, error)
}

package api

import (
	"context"
	"database/sql"
	"games-shelf-api-go/cmd/config"
	"games-shelf-api-go/cmd/models"
	"log"
	"os"
	"time"
)

type Server struct {
	Config config.Config
	Shelf  *models.Shelf
}

func (api *Server) Initialize(cfg config.Config) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDBConnection(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	config.SetConfig(cfg)
	api.Shelf = models.NewShelf(db)

}

func openDBConnection(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil

}

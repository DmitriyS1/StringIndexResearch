package main

import (
	"context"
	"github.com/dmitriys1/StringIndexResearch/internal/db"
	storage "github.com/dmitriys1/StringIndexResearch/internal/store"
	"log"
	"os"
)

func main() {
	cfg := config{
		addr:   os.Getenv("APP_ADDRESS"),
		dbHost: os.Getenv("DB_HOST"),
		dbPort: os.Getenv("DB_PORT"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASSWORD"),
		dbName: os.Getenv("DB_NAME"),
		env:    os.Getenv("ENV"),
	}
	pgxPool := db.NewPostgresDb(context.Background(), cfg.GetDbUrl())
	cmtsStore := storage.NewCommentStore(&pgxPool)
	cndtsStore := storage.NewCandidateStore(&pgxPool)
	store := storage.NewStorage(*cmtsStore, *cndtsStore)

	app := api{
		config: cfg,
		store:  store,
	}

	err := app.run()
	if err != nil {
		log.Fatal(err)
	}
}

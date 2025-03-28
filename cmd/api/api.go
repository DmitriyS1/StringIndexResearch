package main

import (
	"fmt"
	"github.com/dmitriys1/StringIndexResearch/handlers"
	"log"
	"net/http"

	"github.com/dmitriys1/StringIndexResearch/internal/store"
)

type config struct {
	addr   string
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
	env    string
}

func (cfg *config) GetDbUrl() string {
	return "postgres://" + cfg.dbUser + ":" + cfg.dbPass + "@" + cfg.dbHost + ":" + cfg.dbPort + "/" + cfg.dbName + "?sslmode=disable"
}

type app struct {
	config config
	store  *store.Storage
}

func (app *app) run() error {
	fmt.Println("Starting server on " + app.config.addr)
	router := http.NewServeMux()

	commentsHandler := handlers.NewCommentsHandler(app.store)
	candidatesHandler := handlers.NewCandidatesHandler(app.store)

	router.HandleFunc("GET /api/v1/comments/full/{search}", commentsHandler.FullSearchComments)
	router.HandleFunc("GET /api/v1/comments/starts/{search}", commentsHandler.StartsSearchComments)
	router.HandleFunc("GET /api/v1/comments/ends/{search}", commentsHandler.EndsSearchComments)
	router.HandleFunc("GET /api/v1/comments/{id}", commentsHandler.GetById)

	router.HandleFunc("GET /api/v1/candidates/full/{search}", candidatesHandler.FullSearchCandidates)
	router.HandleFunc("GET /api/v1/candidates/starts/{search}", candidatesHandler.StartsSearchCandidates)
	router.HandleFunc("GET /api/v1/candidates/ends/{search}", candidatesHandler.EndsSearchCandidates)
	router.HandleFunc("GET /api/v1/candidates/{id}", candidatesHandler.GetById)

	err := http.ListenAndServe(app.config.addr, router)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

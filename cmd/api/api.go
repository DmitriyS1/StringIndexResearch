package main

import (
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

type app struct {
	config config
	store  *store.Storage
}

func (app *app) run() error {
	router := http.NewServeMux()

	commentsHandler := handlers.NewCommentsHandler(app.store)
	candidatesHandler := handlers.NewCandidatesHandler(app.store)

	router.HandleFunc("GET /api/v1/comments/full/{search}", commentsHandler.FullSearchComments)
	router.HandleFunc("GET /api/v1/comments/starts/{search}", commentsHandler.StartsSearchComments)
	router.HandleFunc("GET /api/v1/comments/ends/{search}", commentsHandler.EndsSearchComments)
	router.HandleFunc("GET /api/v1/comments/{id}", commentsHandler.GetById)
	router.HandleFunc("GET /api/v1/comments/{id}", handler.getCommentHandler)
	router.HandleFunc("GET /api/v1/candidates/titles/{search}", handler.candidatesTitlesHandler)

	err := http.ListenAndServe(app.config.addr, router)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

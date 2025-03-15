package main

import (
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
	store  &store.Store
}

func (app *app) run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/comments/{search}", handler.searchCommentsHandler)
	router.HandleFunc("GET /api/v1/comments/{id}", handler.getCommentHandler)
	router.HandleFunc("GET /api/v1/candidates/titles/{search}", handler.candidatesTitlesHandler)

	err := http.ListenAndServe(app.config.addr, router)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

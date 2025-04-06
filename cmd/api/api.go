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

type api struct {
	config config
	store  *store.Storage
}

func (app *api) run() error {
	fmt.Println("Starting server on " + app.config.addr)
	router := http.NewServeMux()

	candidatesHandler := handlers.NewCandidatesHandler(app.store)

	router.HandleFunc("GET /api/v1/candidates/full/{search}", candidatesHandler.FullSearchCandidates)
	router.HandleFunc("GET /api/v1/candidates/{id}", candidatesHandler.GetById)

	err := http.ListenAndServe(app.config.addr, router)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

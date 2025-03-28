package main

import (
	"context"
	"github.com/dmitriys1/StringIndexResearch/internal/db"
)

func main() {
	pDb := db.NewPostgresDb(context.Background(), "postgres://admin:adminpassword@localhost:5432/string_index?sslmode=disable")

	err := db.SeedCandidatesTable(pDb)
	if err != nil {
		panic(err)
	}
}

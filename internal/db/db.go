package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func NewPostgresDb(dbUrl string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalln("Unable to parse database URL:", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	return db, nil
}

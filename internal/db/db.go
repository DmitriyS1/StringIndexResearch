package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostgresDb struct {
	DB *pgxpool.Pool
}

func NewPostgresDb(ctx context.Context, dbUrl string) PostgresDb {
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalln("Unable to parse database URL:", err)
	}

	db, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	return PostgresDb{DB: db}
}

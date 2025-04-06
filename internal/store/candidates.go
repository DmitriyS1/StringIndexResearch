package store

import (
	"context"
	"fmt"
	"github.com/dmitriys1/StringIndexResearch/internal/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Candidate struct {
	ID        int64  `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	Created   pgtype.Date
}

type CandidatesStore struct {
	db *pgxpool.Pool
}

func NewCandidateStore(db *db.PostgresDb) *CandidatesStore {
	return &CandidatesStore{db: db.DB}
}

func (s *CandidatesStore) FullSearch(ctx context.Context, query string, page int, amount int) ([]Candidate, error) {
	t := time.Now()
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE title ILIKE '%'||$1||'%' LIMIT $2 OFFSET $3", query, amount, page*amount)
	fmt.Printf("Query in FullSearch took: %v with Search string: %s\n", time.Since(t), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidatesStore) StartsWithSearch(ctx context.Context, query string) ([]Candidate, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE title ILIKE $1||'%'", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidatesStore) EndsWithSearch(ctx context.Context, query string) ([]Candidate, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE title ILIKE '%'||$1", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidatesStore) GetById(ctx context.Context, id int64) (*Candidate, error) {
	query := `SELECT * FROM candidates WHERE id = $1`

	var c Candidate

	err := s.db.QueryRow(ctx, query, id).Scan(&c)
	if err != nil {
		return &Candidate{}, err
	}

	return &c, nil
}

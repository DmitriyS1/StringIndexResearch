package store

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Candidate struct {
	ID        int64  `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Title     string `json:"title"`
	Created   pgtype.Date
}

type CandidateStore struct {
	db *pgxpool.Pool
}

func NewCandidateStore(db *pgxpool.Pool) *CandidateStore {
	return &CandidateStore{db: db}
}

func (s *CandidateStore) FullSearch(ctx context.Context, query string) ([]Candidate, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE first_name LIKE $1", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidateStore) StartsWithSearch(ctx context.Context, query string) ([]Candidate, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE first_name LIKE $1%", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidateStore) EndsWithSearch(ctx context.Context, query string) ([]Candidate, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM candidates WHERE first_name LIKE %$1", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[Candidate])
}

func (s *CandidateStore) GetById(ctx context.Context, id int64) (Candidate, error) {
	query := `SELECT * FROM candidates WHERE id = $1`

	var c Candidate

	err := s.db.QueryRow(ctx, query, id).Scan(&c)
	if err != nil {
		return Candidate{}, err
	}

	return c, nil
}

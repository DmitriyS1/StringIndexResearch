package store

import (
	"context"
	"github.com/dmitriys1/StringIndexResearch/internal/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	ID      int64       `json:"id"`
	Text    string      `json:"text"`
	Created pgtype.Date `json:"created"`
}

type CommentsStore struct {
	db *pgxpool.Pool
}

func NewCommentStore(db *db.PostgresDb) *CommentsStore {
	return &CommentsStore{db: db.DB}
}

func (s *CommentsStore) FullSearch(ctx context.Context, query string) ([]Comment, error) {
	rows, err := s.db.Query(ctx, "SELECT * FROM comments WHERE text ILIKE %$1%", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[Comment])
}

func (s *CommentsStore) GetById(ctx context.Context, id int64) (*Comment, error) {
	row := s.db.QueryRow(ctx, "SELECT * FROM comments WHERE id = $1", id)
	var c Comment
	err := row.Scan(&c.ID, &c.Text, &c.Created)
	if err != nil {
		return &Comment{}, err
	}

	return &c, nil
}

package store

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID      int64       `json:"id"`
	Text    string      `json:"text"`
	Created pgtype.Date `json:"created"`
}

type CommentStore struct {
	db *db
}

package store

import "context"

type Storage struct {
	Comments interface {
		Search(context.Context, string) ([]Comment, error)
		GetById(context.Context, int64) (*Comment, error)
	}
	Candidates interface {
		Search(context.Context, string) ([]Candidate, error)
		GetById(context.Context, int64) (*Candidate, error)
	}
}

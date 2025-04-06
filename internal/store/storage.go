package store

import "context"

type Storage struct {
	Comments interface {
		FullSearch(context.Context, string) ([]Comment, error)
		GetById(context.Context, int64) (*Comment, error)
	}
	Candidates interface {
		FullSearch(ctx context.Context, query string, page int, amount int) ([]Candidate, error)
		GetById(context.Context, int64) (*Candidate, error)
	}
}

func NewStorage(comments CommentsStore, candidates CandidatesStore) *Storage {
	return &Storage{
		Comments:   &comments,
		Candidates: &candidates,
	}
}

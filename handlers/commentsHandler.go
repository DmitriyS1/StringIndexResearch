package handlers

import (
	"github.com/dmitriys1/StringIndexResearch/internal/store"
	"net/http"

	hlp "github.com/dmitriys1/StringIndexResearch/helpers/http"
)

type CommentsHandler struct {
	storage *store.Storage
}

func NewCommentsHandler(storage *store.Storage) *CommentsHandler {
	return &CommentsHandler{storage: storage}
}

func (h *CommentsHandler) FullSearchComments(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	comments, err := h.storage.Comments.FullSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(comments, w, r)
	return
}

func (h *CommentsHandler) StartsSearchComments(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	comments, err := h.storage.Comments.StartsWithSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(comments, w, r)
	return
}

func (h *CommentsHandler) EndsSearchComments(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	comments, err := h.storage.Comments.EndsWithSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(comments, w, r)
	return
}

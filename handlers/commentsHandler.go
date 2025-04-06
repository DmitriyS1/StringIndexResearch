package handlers

import (
	"github.com/dmitriys1/StringIndexResearch/internal/store"
	"log"
	"net/http"
	"strconv"

	hlp "github.com/dmitriys1/StringIndexResearch/helpers/http"
)

type CommentsHandler struct {
	storage *store.Storage
}

func NewCommentsHandler(storage *store.Storage) *CommentsHandler {
	return &CommentsHandler{storage: storage}
}

func (h *CommentsHandler) FullSearchComments(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("FullSearchComments")
	search := r.URL.Path[len("/api/v1/comments/full/"):]
	comments, err := h.storage.Comments.FullSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(comments, w, r)
	return
}

func (h *CommentsHandler) GetById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := h.storage.Comments.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(comment, w, r)
	return
}

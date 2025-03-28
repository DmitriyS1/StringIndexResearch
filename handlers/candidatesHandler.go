package handlers

import (
	hlp "github.com/dmitriys1/StringIndexResearch/helpers/http"
	"github.com/dmitriys1/StringIndexResearch/internal/store"
	"net/http"
	"strconv"
)

type CandidatesHandler struct {
	storage *store.Storage
}

func NewCandidatesHandler(storage *store.Storage) *CandidatesHandler {
	return &CandidatesHandler{storage: storage}
}

func (h *CandidatesHandler) FullSearchCandidates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	search := r.URL.Path[len("/api/v1/candidates/full/"):]
	candidates, err := h.storage.Candidates.FullSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(candidates, w, r)
	return
}

func (h *CandidatesHandler) StartsSearchCandidates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	search := r.URL.Query().Get("search")
	candidates, err := h.storage.Candidates.StartsWithSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(candidates, w, r)
	return
}

func (h *CandidatesHandler) EndsSearchCandidates(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	search := r.URL.Query().Get("search")
	candidates, err := h.storage.Candidates.EndsWithSearch(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(candidates, w, r)
	return
}

func (h *CandidatesHandler) GetById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	candidate, err := h.storage.Candidates.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hlp.RespondOk(candidate, w, r)
	return
}

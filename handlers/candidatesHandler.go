package handlers

import (
	"fmt"
	hlp "github.com/dmitriys1/StringIndexResearch/helpers/http"
	"github.com/dmitriys1/StringIndexResearch/internal/store"
	"net/http"
	"strconv"
	"time"
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
	pageParam := r.URL.Query().Get("page")
	amountParam := r.URL.Query().Get("amount")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 0
	}

	amount, err := strconv.Atoi(amountParam)
	if err != nil {
		amount = 100
	}

	t := time.Now()
	candidates, err := h.storage.Candidates.FullSearch(r.Context(), search, page, amount)
	fmt.Printf("Request to DB with mapping took: %v with Search string: %s; Page: %d; Amount: %d\n", time.Since(t), search, page, amount)
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

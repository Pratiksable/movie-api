package handler

import (
	"context"
	"encoding/json"
	"errors"
	"movie-api/internal/model"
	"movie-api/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type MovieStore interface {
	GetAll(context.Context) ([]model.Movie, error)
	GetByID(context.Context, string) (model.Movie, error)
	Create(context.Context, model.Movie) (model.Movie, error)
	Update(context.Context, string, model.Movie) (model.Movie, error)
	Delete(context.Context, string) error
}

type MovieHandler struct{ store MovieStore }

func NewMovieHandler(store MovieStore) *MovieHandler { return &MovieHandler{store: store} }

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies, err := h.store.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Unable to get movies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movie, err := h.store.GetByID(r.Context(), mux.Vars(r)["id"])
	if errors.Is(err, repository.ErrMovieNotFound) {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Unable to get movie", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if movie.Director == nil {
		http.Error(w, "Director is required", http.StatusBadRequest)
		return
	}
	createdMovie, err := h.store.Create(r.Context(), movie)
	if err != nil {
		http.Error(w, "Unable to create movie", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMovie)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if movie.Director == nil {
		http.Error(w, "Director is required", http.StatusBadRequest)
		return
	}
	updatedMovie, err := h.store.Update(r.Context(), mux.Vars(r)["id"], movie)
	if errors.Is(err, repository.ErrMovieNotFound) {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Unable to update movie", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedMovie)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	err := h.store.Delete(r.Context(), mux.Vars(r)["id"])
	if errors.Is(err, repository.ErrMovieNotFound) {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Unable to delete movie", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

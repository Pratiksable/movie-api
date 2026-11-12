package router

import (
	"movie-api/internal/handler"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")

	return r
}

package router

import (
	"movie-api/internal/handler"

	"github.com/gorilla/mux"
)

func New(movieHandler *handler.MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", movieHandler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movieHandler.DeleteMovie).Methods("DELETE")

	return r
}

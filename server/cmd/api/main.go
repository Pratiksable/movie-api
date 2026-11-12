package main

import (
	"context"
	"fmt"
	"log"
	"movie-api/internal/config"
	"movie-api/internal/database"
	"movie-api/internal/handler"
	"movie-api/internal/repository"
	"movie-api/internal/router"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("Unable to load the ENV")
		return
	}
	ctx := context.Background()
	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalln("Error connecting to the database from main", err)
	}
	defer pool.Close()

	if err := database.CreateTables(ctx, pool); err != nil {
		log.Fatalln("Unable to create database tables", err)
	}

	fmt.Println("Database connected")

	movieRepository := repository.NewMovieRepository(pool)
	movieHandler := handler.NewMovieHandler(movieRepository)
	r := router.New(movieHandler)

	address := ":" + cfg.Port
	fmt.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(address, r))
}

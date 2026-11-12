package main

import (
	"fmt"
	"log"
	"movie-api/internal/config"
	"movie-api/internal/router"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("Unable to load the ENV")
		return
	}
	r := router.New()

	address := ":" + cfg.Port
	fmt.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(address, r))
}

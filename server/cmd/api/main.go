package main

import (
	"fmt"
	"log"
	"movie-api/internal/router"
	"net/http"
)

func main() {
	r := router.New()
	fmt.Println("Starting server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

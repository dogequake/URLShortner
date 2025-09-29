package main

import (
	"URLShortner/handler"
	"URLShortner/repository"
	"URLShortner/service"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewMemoryRepo()
	svc := service.NewURLService(repo)
	h := handler.NewHandler(svc)

	http.HandleFunc("/shorten", h.ShortenURL)
	http.HandleFunc("/", h.Redirect)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

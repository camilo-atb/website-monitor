package main

import (
	"config-service/internal/adapters/input/httpi"
	"config-service/internal/adapters/output/persistence"
	"config-service/internal/aplication/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	repo := persistence.NewMemoryRepository()

	service := service.NewSiteService(repo)

	handler := httpi.NewHandler(service)

	r := chi.NewRouter()

	r.Post("/sites", handler.Create)
	r.Get("/sites", handler.List)
	r.Put("/sites/{id}", handler.Update)
	r.Delete("/sites/{id}", handler.Delete)

	log.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

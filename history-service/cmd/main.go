package main

import (
	"history-service/internal/application/service"
	"history-service/internal/domain/adapters/input/httpa"
	"history-service/internal/domain/adapters/output/persistence"
	"net/http"
)

func main() {

	repo := persistence.NewMemoryRepository()

	historyService := service.NewHistoryService(repo)

	handler := httpa.NewHandler(historyService)

	mux := http.NewServeMux()
	mux.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.SaveResult(w, r)
		case http.MethodGet:
			handler.GetResults(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8082", mux)
}

package httpserver

import (
	"encoding/json"
	"net/http"
	"pinger/internal/domain/ports"
)

type handler struct {
	results ports.ResultsPort
}

func NewHandler(results ports.ResultsPort) *handler {
	return &handler{results: results}
}

func (h *handler) GetResults(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	results, err := h.results.GetResults(ctx)

	if err != nil {
		http.Error(w, "error obteniendo resultados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)
}

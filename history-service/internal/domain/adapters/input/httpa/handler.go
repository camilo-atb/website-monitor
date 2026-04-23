package httpa

import (
	"encoding/json"
	"history-service/internal/application/service"
	"history-service/internal/domain/model"
	"net/http"
)

type handler struct {
	service *service.HistoryService // no pasamos la copia sino la misma instancia del servicio, para que el handler pueda usarlo para acceder a los datos; el servicio a su vez tiene una referencia al repositorio, que es donde se almacenan los datos
}

func NewHandler(service *service.HistoryService) *handler {
	return &handler{service: service}
}

func (h *handler) SaveResult(w http.ResponseWriter, r *http.Request) {
	var input model.PingResult

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	err = h.service.SaveResult(input)
	if err != nil {
		http.Error(w, "error saving result", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) GetResults(w http.ResponseWriter, r *http.Request) {
	results, err := h.service.GetResults()
	if err != nil {
		http.Error(w, "error getting results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

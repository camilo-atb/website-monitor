package service

import (
	"history-service/internal/domain/model"
	"history-service/internal/domain/ports"
	"math"
)

type HistoryService struct {
	repo ports.RepositoryPort
}

func NewHistoryService(repo ports.RepositoryPort) *HistoryService {
	return &HistoryService{repo: repo}
}

func (s *HistoryService) SaveResult(result model.PingResult) error {
	return s.repo.Save(result)
}

func (s *HistoryService) GetResults() ([]model.ResultSummary, error) {

	data := s.repo.GetAll()

	var results []model.ResultSummary

	for url, pings := range data {

		total := len(pings)
		if total == 0 {
			continue
		}

		var upCount int

		for _, ping := range pings {
			if ping.Status == "UP" {
				upCount++
			}
		}

		uptime := (float64(upCount) / float64(total)) * 100

		last := pings[total-1]

		results = append(results, model.ResultSummary{
			URL:         url,
			Uptime:      round(uptime, 2),
			TotalChecks: total,
			LastStatus:  last.Status,
			LastChecked: last.CheckedAt,
		})
	}

	return results, nil
}

func round(val float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(val*pow) / pow
}

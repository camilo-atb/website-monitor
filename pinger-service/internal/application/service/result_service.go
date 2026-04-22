package service

import (
	"context"
	"math"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
)

type resultService struct {
	repo ports.ResultRepositoryPort
}

func NewResultService(repo ports.ResultRepositoryPort) ports.ResultsPort {
	return &resultService{
		repo: repo,
	}
}

func (r *resultService) GetResults(ctx context.Context) ([]model.ResultSummary, error) {

	data := r.repo.GetAll()

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
			Uptime:      math.Round(uptime*100) / 100,
			TotalChecks: total,
			LastStatus:  last.Status,
			LastChecked: last.CheckedAt,
		})
	}

	return results, nil
}

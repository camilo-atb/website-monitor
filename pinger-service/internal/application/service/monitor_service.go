package service

import (
	"context"
	"log"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
	"time"
)

type monitor struct {
	configService    ports.ConfigServicePort
	httpClient       ports.HTTPClientPort
	resultRepository ports.ResultRepositoryPort
}

func NewMonitor(configService ports.ConfigServicePort, httpClient ports.HTTPClientPort, resultRepository ports.ResultRepositoryPort) ports.MonitorPort {
	return &monitor{
		configService:    configService,
		httpClient:       httpClient,
		resultRepository: resultRepository,
	}
}

func (m *monitor) Run(ctx context.Context) error {
	sites, err := m.configService.GetSites(ctx)

	log.Printf("sites obtenidos: %d", len(sites))

	if err != nil {
		return err
	}

	if len(sites) == 0 {
		return nil
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	for _, site := range sites {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		statusCode, duration, err := m.httpClient.Get(ctx, site.URL)

		result := model.PingResult{
			URL:          site.URL,
			CheckedAt:    time.Now(),
			ResponseTime: duration,
		}

		if err != nil {
			result.Status = "DOWN"
			result.Error = err.Error()
		} else {
			result.StatusCode = statusCode

			if statusCode >= 200 && statusCode < 300 {
				result.Status = "UP"
			} else {
				result.Status = "DOWN"
			}
		}

		log.Printf(
			"[PING] %s | status=%s | code=%d | time=%v | error=%s",
			result.URL,
			result.Status,
			result.StatusCode,
			result.ResponseTime,
			result.Error,
		)

		_ = m.resultRepository.Save(ctx, result) // * No paramos el loop
	}

	return nil
}

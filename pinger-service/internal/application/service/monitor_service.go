package service

import (
	"context"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
	"sync"
	"time"
)

type monitor struct {
	configService  ports.ConfigServicePort
	httpClient     ports.HTTPClientPort
	historyService ports.HistoryPort
}

func NewMonitor(configService ports.ConfigServicePort, httpClient ports.HTTPClientPort, historyService ports.HistoryPort) *monitor {
	return &monitor{
		configService:  configService,
		httpClient:     httpClient,
		historyService: historyService,
	}
}

func (m *monitor) Run(ctx context.Context) error {

	sites, err := m.configService.GetSites(ctx)
	if err != nil {
		return err
	}

	workerCount := 5
	jobs := make(chan model.MonitoredURL)

	var wg sync.WaitGroup

	// workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for site := range jobs {

				select {
				case <-ctx.Done():
					return
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

				_ = m.historyService.Save(ctx, result)
			}
		}()
	}

LOOP: // label que identifica el loop para poder salir de él desde el select
	for _, site := range sites {
		select {
		case <-ctx.Done():
			break LOOP
		case jobs <- site:
		}
	}

	close(jobs)

	wg.Wait()

	return nil
}

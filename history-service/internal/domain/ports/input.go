package ports

import "history-service/internal/domain/model"

type HistoryPort interface {
	SaveResult(result model.PingResult) error
	GetResults() ([]model.ResultSummary, error)
}

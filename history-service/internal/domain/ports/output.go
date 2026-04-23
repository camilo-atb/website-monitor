package ports

import "history-service/internal/domain/model"

type RepositoryPort interface {
	Save(result model.PingResult) error
	GetAll() map[string][]model.PingResult
}

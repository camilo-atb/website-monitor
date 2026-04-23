package persistence

import (
	"history-service/internal/domain/model"
	"history-service/internal/domain/ports"
	"sync"
)

type memoryRepository struct {
	mu      sync.Mutex
	results map[string][]model.PingResult // los maps en go no son thread-safe
}

func NewMemoryRepository() ports.RepositoryPort {
	return &memoryRepository{
		results: make(map[string][]model.PingResult),
	}
}

func (m *memoryRepository) Save(result model.PingResult) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.results[result.URL] = append(m.results[result.URL], result)
	return nil
}

func (m *memoryRepository) GetAll() map[string][]model.PingResult {
	m.mu.Lock()
	defer m.mu.Unlock()

	// copia defensiva para evitar que el caller modifique el estado interno del repositorio; no devolvemos el map original
	copyMap := make(map[string][]model.PingResult)

	for k, v := range m.results {
		copyMap[k] = append([]model.PingResult{}, v...)
	}

	return copyMap
}

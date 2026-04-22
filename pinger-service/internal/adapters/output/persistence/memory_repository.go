package persistence

import (
	"context"
	"pinger/internal/domain/model"
	"pinger/internal/domain/ports"
	"sync"
)

type memoryRepository struct {
	mu      sync.Mutex
	results map[string][]model.PingResult
}

func NewMemoryRepository() ports.ResultRepositoryPort {
	return &memoryRepository{
		results: make(map[string][]model.PingResult),
	}
}

func (r *memoryRepository) Save(ctx context.Context, result model.PingResult) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	r.mu.Lock() // * para que no colicionen los goroutine que intenten guardar resultados al mismo tiempo, ya que el map no es seguro para concurrencia
	defer r.mu.Unlock()

	r.results[result.URL] = append(r.results[result.URL], result)

	return nil
}

func (r *memoryRepository) GetByURL(url string) ([]model.PingResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	results, ok := r.results[url]
	if !ok {
		return []model.PingResult{}, nil
	}

	return results, nil
}

func (r *memoryRepository) GetAll() map[string][]model.PingResult {
	r.mu.Lock()
	defer r.mu.Unlock()

	copyMap := make(map[string][]model.PingResult)

	for k, v := range r.results { // Encapsulación de estado mutable
		copyMap[k] = append([]model.PingResult{}, v...)
	}

	return copyMap

	// ! return r.results // No es bueno devolver el map original porque podría ser modificado por el caller
}

// * Un código es thread-safe (seguro para hilos) cuando: Puede ser usado por múltiples hilos o goroutines al mismo tiempo sin causar errores o comportamientos inesperados.

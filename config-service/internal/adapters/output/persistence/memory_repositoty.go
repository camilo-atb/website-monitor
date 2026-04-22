package persistence

import (
	"config-service/internal/domain/model"
	"config-service/internal/domain/ports"
	"errors"
	"sync"
)

type memoryRepository struct {
	data   map[int]model.MonitoredURL
	lastID int
	mu     sync.Mutex
}

func NewMemoryRepository() ports.OutputPort {
	return &memoryRepository{
		data: make(map[int]model.MonitoredURL),
	}
}

func (r *memoryRepository) Save(site model.MonitoredURL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	site.ID = r.lastID

	r.data[site.ID] = site
	return nil
}

func (r *memoryRepository) FindByID(id int) (model.MonitoredURL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	site, ok := r.data[id]
	if !ok {
		return model.MonitoredURL{}, errors.New("site no encontrado")
	}

	return site, nil
}

func (r *memoryRepository) FindAll() ([]model.MonitoredURL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var sites []model.MonitoredURL
	for _, site := range r.data {
		sites = append(sites, site)
	}

	return sites, nil
}

func (r *memoryRepository) Update(site model.MonitoredURL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[site.ID]
	if !ok {
		return errors.New("site no encontrado")
	}

	r.data[site.ID] = site
	return nil
}

func (r *memoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.data[id]
	if !ok {
		return errors.New("site no encontrado")
	}

	delete(r.data, id)
	return nil
}

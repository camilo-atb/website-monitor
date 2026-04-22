package ports

import "config-service/internal/domain/model"

type OutputPort interface {
	Save(model.MonitoredURL) error
	FindByID(id int) (model.MonitoredURL, error)
	FindAll() ([]model.MonitoredURL, error)
	Update(model.MonitoredURL) error
	Delete(id int) error
}

/*
Output port:
Guardar en DB
Obtener por ID
Actualizar
Eliminar
Ver todos
*/

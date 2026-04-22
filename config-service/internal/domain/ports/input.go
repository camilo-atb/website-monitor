package ports

import (
	"config-service/internal/domain/model"
	"time"
)

// Lo que el sistema ofrece
type CreateSiteInput struct {
	URL        string        `json:"url"`
	ReviewTime time.Duration `json:"reviewTime"`
}

type UpdateSiteInput struct {
	URL        *string        `json:"url"`
	ReviewTime *time.Duration `json:"reviewTime"`
}

type InputPort interface {
	Create(input CreateSiteInput) error
	Update(ID int, input UpdateSiteInput) error
	Delete(ID int) error
	List() ([]model.MonitoredURL, error)
}

/*
Input port:
Crear sitio
Actualizar
Eliminar
Listar
*/

/*
¿Por qué es mejor el CreateSiteInput que pasar los datos sueltos?
Porque es más fácil de mantener, si en el futuro se agregan más campos, solo se modifica el struct y no todas las funciones que lo usan.

Ademas, separas:
datos de entrada
modelo de dominio
👉 Esto evita acoplamiento innecesario
*/

package ports

import (
	"context"
	"pinger/internal/domain/model"
	"time"
)

// * El service no sabe hacer todo solo necesita de otros servicios para hacer su trabajo, por eso se definen estas interfaces que el service va a usar para interactuar con el mundo exterior, como por ejemplo hacer peticiones HTTP, acceder a la base de datos, etc.

type HTTPClientPort interface {
	Get(ctx context.Context, url string) (int, time.Duration, error)
} // * Es el que hace el ping; practicamente el que hace la petición HTTP

type HistoryPort interface {
	Save(ctx context.Context, result model.PingResult) error
}

type ConfigServicePort interface { // * Es el puente hacia el otro microservicio
	GetSites(ctx context.Context) ([]model.MonitoredURL, error)
} // * El que obtiene la configuración de los sitios a monitorear, en este caso de la base de datos, pero podría ser de un archivo, una API, etc. Ya que, el pinger no sabe qué URLs existen o cómo obtenerlas, solo sabe que tiene que obtenerlas de alguna parte, y esta interfaz le dice cómo obtenerlas. Estas URLs las obtenemos del service de configuración (config-service)

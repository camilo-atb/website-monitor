package ports

import (
	"context"
	"pinger/internal/domain/model"
)

// * InputPort → lo que el sistema hace

type MonitorPort interface {
	// Run guarda resultados
	Run(ctx context.Context) error // * Esto es un proceso interno, no es algo que se exponga a través de una API, sino que es algo que se ejecuta de forma periódica para hacer el monitoreo. Es el que le dice al sistema que ejecute la lógica de monitoreo.
} // * Es el que le dice al sistema que ejecute la lógica de monitoreo

type ResultsPort interface {
	// GetResults calcula métricas
	GetResults(ctx context.Context) ([]model.ResultSummary, error) // * caso de uso externo, es decir, algo que se expone a través de una API para que otros sistemas puedan obtener los resultados del monitoreo. Es el que le dice al sistema que devuelva los resultados del monitoreo.
}

/*
CQRS
CQRS significa Command Query Responsibility Segregation (en español, Segregación de Responsabilidades entre Comandos y Consultas). Es un patrón de arquitectura en el desarrollo de software.

Commands: es una operación que modifica el estado del sistema, como crear, actualizar o eliminar datos. Por ejemplo, en un sistema de gestión de usuarios, un comando podría ser "CrearUsuario" o "ActualizarUsuario".

Queries: es una operación que consulta el estado del sistema sin modificarlo, como obtener datos o realizar búsquedas. Por ejemplo, en el mismo sistema de gestión de usuarios, una consulta podría ser "ObtenerUsuarioPorID" o "ListarUsuarios".

En CQRS, se separan las responsabilidades de los comandos y las consultas en diferentes modelos o componentes. Esto permite optimizar cada uno de ellos para su propósito específico. Por ejemplo, el modelo de comandos puede estar diseñado para manejar transacciones y garantizar la consistencia de los datos, mientras que el modelo de consultas puede estar optimizado para la lectura y el rendimiento.
*/

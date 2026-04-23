package ports

import (
	"context"
)

// * InputPort → lo que el sistema hace

type MonitorPort interface {
	// Run guarda resultados
	Run(ctx context.Context) error // * Esto es un proceso interno, no es algo que se exponga a través de una API, sino que es algo que se ejecuta de forma periódica para hacer el monitoreo. Es el que le dice al sistema que ejecute la lógica de monitoreo.
} // * Es el que le dice al sistema que ejecute la lógica de monitoreo

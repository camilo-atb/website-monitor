package ports

import "context"

// * InputPort → lo que el sistema hace

type MonitorPort interface {
	Run(ctx context.Context) error
} // * Es el que le dice al sistema que ejecute la lógica de monitoreo

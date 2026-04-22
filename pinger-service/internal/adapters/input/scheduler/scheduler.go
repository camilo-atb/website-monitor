package scheduler

import (
	"context"
	"log"
	"pinger/internal/domain/ports"
	"time"
)

type scheduler struct { // Input adapter: este componente es quien dispara la lógica del dominio.
	monitor  ports.MonitorPort // la tarea a ejecutar (interfaz)
	interval time.Duration     // cada cuánto tiempo ejecutarla
}

func NewScheduler(monitor ports.MonitorPort, interval time.Duration) *scheduler {
	return &scheduler{
		monitor:  monitor,
		interval: interval,
	}
}

func (s *scheduler) Start(ctx context.Context) {
	log.Println("ejecutando ciclo de monitoreo...")
	for {
		select {
		case <-ctx.Done():
			log.Println("scheduler detenido")
			return

		default:
			err := s.monitor.Run(ctx)
			if err != nil {
				log.Println("error running monitor:", err)
			}

			select {
			case <-ctx.Done():
				log.Println("scheduler detenido")
				return
			case <-time.After(s.interval):
			}
		}
	}
}

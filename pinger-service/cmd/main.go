package main

import (
	"context"
	"os/signal"
	"pinger/internal/adapters/input/scheduler"
	"pinger/internal/adapters/output/configclient"
	"pinger/internal/adapters/output/httpclient"
	"pinger/internal/adapters/output/persistence"
	"pinger/internal/application/service"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	repo := persistence.NewMemoryRepository()
	httpClient := httpclient.NewHTTPClient()
	configClient := configclient.NewConfigClient()

	service := service.NewMonitor(configClient, httpClient, repo)

	scheduler := scheduler.NewScheduler(service, 10*time.Second)
	scheduler.Start(ctx)
}

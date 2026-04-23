package main

import (
	"context"
	"os/signal"
	"pinger/internal/adapters/input/scheduler"
	"pinger/internal/adapters/output/configclient"
	"pinger/internal/adapters/output/historyclient"
	"pinger/internal/adapters/output/httpclient"
	"pinger/internal/application/service"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpClient := httpclient.NewHTTPClient()
	configClient := configclient.NewConfigClient()
	historyClient := historyclient.NewHistoryClient("http://localhost:8082")

	monitorService := service.NewMonitor(configClient, httpClient, historyClient)

	scheduler := scheduler.NewScheduler(monitorService, 10*time.Second)
	go scheduler.Start(ctx)

	<-ctx.Done()
	println("apagando sistema...")

	println("servidor detenido correctamente")
}

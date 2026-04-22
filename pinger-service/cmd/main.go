package main

import (
	"context"
	"net/http"
	"os/signal"
	"pinger/internal/adapters/input/httpserver"
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

	servicee := service.NewMonitor(configClient, httpClient, repo)
	resultService := service.NewResultService(repo)

	handler := httpserver.NewHandler(resultService)

	scheduler := scheduler.NewScheduler(servicee, 10*time.Second)
	go scheduler.Start(ctx) // sheduler es bloqueante. Por eso lo corremos en una goroutine

	mux := http.NewServeMux()
	mux.HandleFunc("/results", handler.GetResults)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

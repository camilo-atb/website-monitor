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

	monitorService := service.NewMonitor(configClient, httpClient, repo)
	resultService := service.NewResultService(repo)

	handler := httpserver.NewHandler(resultService)

	scheduler := scheduler.NewScheduler(monitorService, 10*time.Second)
	go scheduler.Start(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/results", handler.GetResults)

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Para el Ctrl + C
	<-ctx.Done()
	println("apagando sistema...")

	// Apagar servidor http
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		panic(err)
	}

	println("servidor detenido correctamente")
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"ride-hail/internal/config"
	"ride-hail/internal/pkg/log"
	"ride-hail/internal/shutdown"
)

func main() {
	logger := log.New("ride-service")
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("cfg_load", "failed to load config", "req-123", "ride-456", err)
		return
	}

	addr := fmt.Sprintf(":%d", cfg.Services.RideService)
	server := &http.Server{
		Addr:    addr,
		Handler: http.DefaultServeMux,
	}

	go func() {
		logger.Info("server_start", fmt.Sprintf("Listening on %s", addr), "req-id", "", "port", cfg.Services.RideService)
		server.ListenAndServe()
	}()

	// 👇 Вместо кучи кода
	shutdown.GracefulStop(server.Shutdown, 10*time.Second)
	logger.Info("shutdown", "Server stopped gracefully ✅", "req-id", "", "port", cfg.Services.RideService)
}

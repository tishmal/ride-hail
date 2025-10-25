package main

import (
	"fmt"
	"net/http"
	"time"

	"ride-hail/internal/config"
	"ride-hail/internal/db"
	"ride-hail/internal/microservices/ride"
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

	conn, err := db.ConnectPostgres(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)
	if err != nil {
		logger.Error("db_connect", "failed to connect to PostgreSQL", "req-id", "ride-456", err)
		return
	}
	defer conn.Close()

	logger.Info("db_connect", "PostgreSQL connection established", "req-id", "ride-456")

	rideServer := ride.NewServer(conn)
	handler := rideServer.Routes()

	addr := fmt.Sprintf(":%d", cfg.Services.RideService)
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		logger.Info("server_start", fmt.Sprintf("Listening on %s", addr), "req-id", "", "port", cfg.Services.RideService)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server_error", "HTTP server error", "req-id", "", err)
		}
	}()

	shutdown.GracefulStop(server.Shutdown, 10*time.Second)
	logger.Info("shutdown", "Server stopped gracefully ✅", "req-id", "", "port", cfg.Services.RideService)
}

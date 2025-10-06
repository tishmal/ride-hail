package main

import (
	"errors"
	"fmt"
	"ride-hail/internal/config"
	log "ride-hail/internal/pkg/logger"
)

func main() {
	logger := log.New("ride-service")
	logger.Info("startup", "ride svc started", "req1", "", "port", 3)

	cfg, err := config.LoadConfig()
	if err != nil {
		err := errors.New("load config failed")
		logger.Error("cfg_load", "", "req-123", "ride-456", err)
	}

	fmt.Printf("Ride service is running on port %d 🚖\n", cfg.Services.RideService)
}

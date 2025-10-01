package main

import (
	"fmt"
	"ride-hail/internal/config"
	log "ride-hail/internal/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.InitLogger("ride-service")
	log.Logger.Info("service started", "port", cfg.Services.RideService)

	fmt.Println(cfg)
	fmt.Println(cfg.Services.RideService)
	fmt.Println("Ride service is running 🚖")
}

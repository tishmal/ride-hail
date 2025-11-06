package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ridehttp "ride-hail/internal/microservices/ride/http"
	"ride-hail/internal/shared/config"
	"ride-hail/internal/shared/logger"
	"ride-hail/internal/shared/mq"
)

func main() {
	log := logger.New("Ride")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("config", "failed to load config", "err", "", err)
		panic(err)
	}

	conn, ch, err := mq.ConnectRabbit(cfg)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	defer ch.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/rides", ridehttp.HandleCreateRide(ch))

	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.Services.RideService), Handler: mux}
	go func() {
		log.Info("ride-service", "INFO", "start", "Ride service started")
		_ = srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	srv.Close()
	log.Info("ride-service", "INFO", "shutdown", "Ride service stopped")
	select {}
}

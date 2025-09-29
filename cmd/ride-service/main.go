package main

import (
	"fmt"
	"log"
	"ride-hail/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(cfg)
	fmt.Println(cfg.Services.RideService)
	fmt.Println("Ride service is running 🚖")
}

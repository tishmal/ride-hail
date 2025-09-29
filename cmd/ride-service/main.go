package main

import (
	"fmt"
	"ride-hail/internal/config"
	"strconv"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg.App.Port)
	port, err := strconv.Atoi(cfg.App.Port)

	if err != nil {
		fmt.Println("GG")
	}
	fmt.Println("Ride service is running 🚖")
	fmt.Println(port)
}

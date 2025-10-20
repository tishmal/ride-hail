package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	RabbitMQ struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`
	WebSocket struct {
		Port int `yaml:"port"`
	} `yaml:"websocket"`
	Services struct {
		RideService           int `yaml:"ride_service"`
		DriverLocationService int `yaml:"driver_location_service"`
		AdminService          int `yaml:"admin_service"`
	} `yaml:"services"`
}

func LoadConfig() (*Config, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "app/config.yaml"
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file at %s: %v", path, err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	CheckEmptyCfg(&cfg)

	return &cfg, nil
}

func CheckEmptyCfg(cfg *Config) {
	if cfg.Database.Host == "" {
		cfg.Database.Host = os.Getenv("DATABASE_HOST")
	}
	if cfg.Database.Port == 0 {
		if portStr := os.Getenv("DATABASE_PORT"); portStr != "" {
			cfg.Database.Port, _ = strconv.Atoi(portStr)
		}
	}
	if cfg.Database.User == "" {
		cfg.Database.User = os.Getenv("DB_USER")
	}
	if cfg.Database.Password == "" {
		cfg.Database.Password = os.Getenv("DB_PASSWORD")
	}
	if cfg.Database.Name == "" {
		cfg.Database.Name = os.Getenv("DB_NAME")
	}

	if cfg.RabbitMQ.Host == "" {
		cfg.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	}
	if cfg.RabbitMQ.Port == 0 {
		if portStr := os.Getenv("RABBITMQ_PORT"); portStr != "" {
			cfg.RabbitMQ.Port, _ = strconv.Atoi(portStr)
		}
	}
	if cfg.RabbitMQ.User == "" {
		cfg.RabbitMQ.User = os.Getenv("RABBITMQ_USER")
	}
	if cfg.RabbitMQ.Password == "" {
		cfg.RabbitMQ.Password = os.Getenv("RABBITMQ_PASSWORD")
	}

	if cfg.WebSocket.Port == 0 {
		if portStr := os.Getenv("WS_PORT"); portStr != "" {
			cfg.WebSocket.Port, _ = strconv.Atoi(portStr)
		}
	}

	if cfg.Services.RideService == 0 {
		if portStr := os.Getenv("RIDE_SERVICE_PORT"); portStr != "" {
			cfg.Services.RideService, _ = strconv.Atoi(portStr)
		}
	}
	if cfg.Services.DriverLocationService == 0 {
		if portStr := os.Getenv("DRIVER_LOCATION_SERVICE_PORT"); portStr != "" {
			cfg.Services.DriverLocationService, _ = strconv.Atoi(portStr)
		}
	}
	if cfg.Services.AdminService == 0 {
		if portStr := os.Getenv("ADMIN_SERVICE_PORT"); portStr != "" {
			cfg.Services.AdminService, _ = strconv.Atoi(portStr)
		}
	}
}

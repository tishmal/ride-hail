package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"database"`
	} `yaml:"database"`

	RabbitMQ struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`

	WebSocket struct {
		Port string `yaml:"port"`
	} `yaml:"websocket"`

	Services struct {
		RideService           string `yaml:"ride_service"`
		DriverLocationService string `yaml:"driver_location_service"`
		AdminService          string `yaml:"admin_service"`
	} `yaml:"services"`
}

func LoadConfig() (*Config, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "app/config.yaml"
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file at %s: %v", path, err)
	}

	raw = []byte(os.ExpandEnv(string(raw)))

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}

	FillFromEnv(&cfg)
	return &cfg, nil
}

func FillFromEnv(cfg *Config) {
	// Database
	if cfg.Database.Host == "" {
		cfg.Database.Host = os.Getenv("DB_HOST")
	}
	if cfg.Database.Port == "" {
		cfg.Database.Port = os.Getenv("DB_PORT")
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

	// Rabbit
	if cfg.RabbitMQ.Host == "" {
		cfg.RabbitMQ.Host = os.Getenv("RABBITMQ_HOST")
	}
	if cfg.RabbitMQ.Port == "" {
		cfg.RabbitMQ.Port = os.Getenv("RABBITMQ_PORT")
	}
	if cfg.RabbitMQ.User == "" {
		cfg.RabbitMQ.User = os.Getenv("RABBITMQ_USER")
	}
	if cfg.RabbitMQ.Password == "" {
		cfg.RabbitMQ.Password = os.Getenv("RABBITMQ_PASSWORD")
	}

	// WebSocket
	if cfg.WebSocket.Port == "" {
		cfg.WebSocket.Port = os.Getenv("WS_PORT")
	}

	// Services ports
	if cfg.Services.RideService == "" {
		cfg.Services.RideService = os.Getenv("RIDE_SERVICE_PORT")
	}
	if cfg.Services.DriverLocationService == "" {
		cfg.Services.DriverLocationService = os.Getenv("DRIVER_LOCATION_SERVICE_PORT")
	}
	if cfg.Services.AdminService == "" {
		cfg.Services.AdminService = os.Getenv("ADMIN_SERVICE_PORT")
	}
}

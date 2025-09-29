package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Port string `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	RabbitMQ struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"rabbitmq"`
}

func LoadConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./config/config.yaml" // или "./internal/config/config.yaml"
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	return &cfg
}

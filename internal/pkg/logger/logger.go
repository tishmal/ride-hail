package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(serviceName string) {
	level := slog.LevelInfo
	if os.Getenv("LOG_LEVEL") == "debug" {
		level = slog.LevelDebug
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	Logger = slog.New(handler).With(
		"service", "ride-service",
		"env", os.Getenv("APP_ENV"),
		"version", "1.0.0",
	)

}

package log

import (
	"log/slog"
	"os"
	"runtime/debug"
	"time"
)

type Logger struct {
	logger  *slog.Logger
	service string
	host    string
}

func New(service string) *Logger {
	host, _ := os.Hostname()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &Logger{
		logger:  slog.New(handler),
		service: service,
		host:    host,
	}
}

func (l *Logger) Info(action, msg, requestID, rideID string, extra ...any) {
	args := []any{
		"timestamp", time.Now().Format(time.RFC3339),
		"level", "INFO",
		"service", l.service,
		"hostname", l.host,
		"action", action,
		"request_id", requestID,
		"ride_id", rideID,
	}
	args = append(args, extra...)

	l.logger.Info(msg, args...)
}

func (l *Logger) Debug(action, msg, requestID, rideID string, extra ...any) {
	args := []any{
		"timestamp", time.Now().Format(time.RFC3339),
		"level", "DEBUG",
		"service", l.service,
		"hostname", l.host,
		"action", action,
		"request_id", requestID,
		"ride_id", rideID,
	}

	args = append(args, extra...)

	l.logger.Debug(msg, args...)
}

func (l *Logger) Error(action, msg, requestID, rideID string, err error, extra ...any) {
	args := []any{
		"timestamp", time.Now().Format(time.RFC3339),
		"level", "ERROR",
		"service", l.service,
		"hostname", l.host,
		"action", action,
		"request_id", requestID,
		"ride_id", rideID,
		"error", map[string]any{
			"msg":   err.Error(),
			"stack": string(debug.Stack()),
		},
	}

	args = append(args, extra...)

	l.logger.Debug(msg, args...)
}

package logger

import (
	"log/slog"
)

var Log *AppLogger

type AppLogger struct {
}

func New() *AppLogger {
	Log = &AppLogger{}
	return Log
}

func (_l *AppLogger) Info(message string) {
	slog.Info(message)
}

func (_l *AppLogger) Warning(message string, details string) {
	slog.Warn(message, "details", details)
}

func (_l *AppLogger) Error(message string, details string, err error) {
	slog.Error(message, "details", details, "error", err.Error())
}

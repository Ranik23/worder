package logger

import (
	"Worder/internal/config"
	"log/slog"
	"os"
	"github.com/lmittmann/tint"
)

func SetUpLogger(cfg *config.Config) *slog.Logger {
	levelLog := cfg.App.LogLevel

	var leveler slog.Level

	switch levelLog {
	case "debug":
		leveler = slog.LevelDebug
	case "warn":
		leveler = slog.LevelWarn
	case "info":
		leveler = slog.LevelInfo
	case "error":
		leveler = slog.LevelError
	default:
		leveler = slog.LevelInfo
	}

	return slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level: leveler,
	}))
}
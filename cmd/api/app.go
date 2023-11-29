package api

import (
	"log/slog"
)

type AppConfig struct {
	Addr   string
	Env    string
	Logger slog.Logger
}

type App interface {
	StartServer() error
}
type app struct {
	config *AppConfig
	logger *slog.Logger
}

func NewApp(config *AppConfig, logger *slog.Logger) App {
	return &app{
		config: config,
		logger: logger,
	}
}

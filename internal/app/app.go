package app

import (
	grpcApp "Worder/internal/app/grpc"
	"Worder/internal/config"
	"Worder/internal/storage"
	"log/slog"
)




type App struct {
	GRPCApp *grpcApp.App
}

func NewApp(logger *slog.Logger, config *config.Config, strg storage.Storage) *App {
	grpcApplication := grpcApp.NewApp(logger, config, strg) 
	return &App{
		GRPCApp: grpcApplication,
	}
}
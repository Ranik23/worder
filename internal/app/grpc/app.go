package grpcApp

import (
	grpcserver "Worder/internal/api"
	"Worder/internal/config"
	"Worder/internal/services"
	corrector_service "Worder/internal/services/corrector"
	"Worder/internal/storage"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
    grpcServer *grpc.Server
    logger     *slog.Logger
    config     *config.Config
    storage    storage.Storage
    services   *services.AppServices
}


func NewApp(logger *slog.Logger, config *config.Config, strg storage.Storage) *App {

    correctorService := corrector_service.NewCorrectorService(strg, logger, config)

    appServices := &services.AppServices{
        Corrector: correctorService,
    }

    grpcServer := grpc.NewServer()

    grpcserver.Register(grpcServer, correctorService)

    return &App{
        grpcServer: grpcServer,
        logger:     logger,
        config:     config,
        storage:    strg,
        services:   appServices,
    }
}

func (a *App) MustRun() {
    if err := a.Run(); err != nil {
        panic(err)
    }
}

func (a *App) Run() error {
    const op = "grpcapp.Run"

    host := a.config.App.Host
    if host == "" {
        host = "0.0.0.0"
    }

    port := a.config.App.Port
    if port == "" {
        port = "50051"
    }

    l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
    if err != nil {
        return fmt.Errorf("%s: failed to listen: %w", op, err)
    }

    defer func() {
        if cerr := l.Close(); cerr != nil {
            a.logger.Error("failed to close listener", slog.String("error", cerr.Error()))
        }
    }()

    a.logger.Info("gRPC server started", slog.String("addr", l.Addr().String()))

    if err := a.grpcServer.Serve(l); err != nil {
        return fmt.Errorf("%s: failed to serve: %w", op, err)
    }

    return nil
}

func (a *App) Stop() {
    const op = "grpcapp.Stop"

    a.logger.With(slog.String("op", op)).
        Info("stopping gRPC server")

    a.grpcServer.GracefulStop()
}

package main

import (
	"Worder/internal/app"
	"Worder/internal/config"
	db "Worder/internal/storage/postgres"
	"Worder/internal/utils/logger"
)



func main() {

	configPath := "/home/anton/Worder/config/config.yaml"

	Config, err := config.LoadConfig(configPath)
	if err != nil {
		return
	}

	Logger := logger.SetUpLogger(Config)

	Storage := db.NewPostgresStorage(Logger, Config)

	App := app.NewApp(Logger, Config, Storage)

	App.GRPCApp.MustRun()

}
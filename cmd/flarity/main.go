package main

import (
	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/server"
	"github.com/jenish-jain/logger"
)

func main() {

	configStore := config.InitConfig("production")
	logger.Init(configStore.LogLevel)
	logger.Info("Starting flarity app")
	server := server.NewServer()
	server.Run()
}

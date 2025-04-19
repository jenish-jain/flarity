package main

import (
	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/ingestor"
	"github.com/jenish-jain/flarity/internal/login"
	"github.com/jenish-jain/flarity/internal/server"
	"github.com/jenish-jain/logger"
)

func main() {

	// ctx := context.Background()
	configStore := config.InitConfig("production")
	logger.Init(configStore.LogLevel)
	logger.Info("Starting flarity app")

	// mongoClient := mongo.NewMongoClient(configStore.MongoConfig, configStore.AppName)
	// defer mongo.Disconnnect(mongoClient, ctx)

	ingestorHandler := ingestor.NewHandler()
	loginHandler := login.NewHandler()
	handlers := server.InitHandlers(ingestorHandler, loginHandler)
	server := server.NewServer()
	server.Run(handlers)
}

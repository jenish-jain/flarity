package main

import (
	"context"

	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/ingestor"
	"github.com/jenish-jain/flarity/internal/login"
	"github.com/jenish-jain/flarity/internal/mongo"
	"github.com/jenish-jain/flarity/internal/server"
	"github.com/jenish-jain/flarity/internal/transaction"
	"github.com/jenish-jain/logger"
)

func main() {

	// ctx := context.Background()
	configStore := config.InitConfig("production")
	logger.Init(configStore.LogLevel)
	logger.Info("Starting flarity app")

	ctx := context.Background()
	mongoClient := mongo.NewMongoClient(configStore, configStore.AppName)
	defer mongo.Disconnect(mongoClient, ctx)
	transactionRepo := transaction.NewRepository(mongoClient, configStore)

	ingestorHandler := ingestor.NewHandler(transactionRepo)
	loginHandler := login.NewHandler()
	handlers := server.InitHandlers(ingestorHandler, loginHandler)
	server := server.NewServer()
	server.Run(handlers)

}

package mongo

import (
	"context"

	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(config *config.Config, appName string) *mongo.Client {
	logger.Info("Initializing mongo client", "appName", appName)

	clientOptions := options.Client().ApplyURI(config.GetMongoURI())
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Error("Error connecting to mongo",
			"event", "MONGO_CLIENT_INITIALISATION",
			"err", err)
		panic(err) // Uncommenting the panic to handle connection errors
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		logger.Error("Error pinging mongo",
			"event", "MONGO_CLIENT_PING",
			"err", err)
		panic(err)
	}

	logger.Info("Mongo client initialized successfully", "appName", appName)

	return client
}

func Disconnect(client *mongo.Client, ctx context.Context) error {
	if err := client.Disconnect(ctx); err != nil {
		logger.Error("Error disconnecting from mongo",
			"event", "MONGO_CLIENT_DISCONNECTION",
			"err", err)
		return err
	}
	logger.Info("Successfully disconnected from mongo", "event", "MONGO_CLIENT_DISCONNECTION")
	return nil
}

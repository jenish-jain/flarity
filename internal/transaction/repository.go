package transaction

import (
	"context"

	"github.com/jenish-jain/flarity/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Add(transaction *Transaction) error
	Get(page, pageSize int) ([]Transaction, error)
}

type repository struct {
	transactionCollection *mongo.Collection
}

func (r *repository) Add(transaction *Transaction) error {
	// Implementation for creating a transaction
	_, err := r.transactionCollection.InsertOne(context.TODO(), transaction)
	return err
}

// update this for paginated query
func (r *repository) Get(page, pageSize int) ([]Transaction, error) {
	// Implementation for getting all transactions with pagination
	skip := (page - 1) * pageSize
	cursor, err := r.transactionCollection.Find(context.TODO(), bson.D{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var transactions []Transaction
	for cursor.Next(context.TODO()) {
		var transaction Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func NewRepository(mongoClient *mongo.Client, config *config.Config) Repository {
	collectionName := "transactions"
	command := bson.D{{Key: "create", Value: collectionName}}
	var result bson.M
	if err := mongoClient.Database(config.GetMongoDbName()).RunCommand(context.TODO(), command).Decode(&result); err != nil {
		panic(err)
	}

	collection := mongoClient.Database(config.GetMongoDbName()).Collection(collectionName)
	return &repository{
		transactionCollection: collection,
	}

}

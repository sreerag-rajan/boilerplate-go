package databaseimplementation

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	URI string
}

func (m *MongoDB) Connect(ctx context.Context) (interface{}, error) {
	clientOptions := options.Client().ApplyURI(m.URI)
	Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return Client, nil
}

func NewMongoDB(connectionString string) *MongoDB {
	return &MongoDB{
		URI: connectionString,
	}
}

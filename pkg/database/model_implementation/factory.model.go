package modelimplementation

import (
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	FindById(id string) (interface{}, error)

	FindOne(criteria interface{}) (interface{}, error)

	FindAll(criteria interface{}) ([]interface{}, error)

	Create(data interface{}) error

	UpdateOne(criteria interface{}, update interface{}) error

	UpdateById(id string, update interface{}) error

	UpdateAll(criteria interface{}, update interface{}) error

	DeleteById(id string) error

	DeleteOne(criteria interface{}) error

	DeleteAll(criteria interface{}) error
}

// ModelFactory is a factory method that returns a ModelImplementation based on the type passed in
func ModelFactory(modelType string, client interface{}, dataSourceName string, dbName string) (Model, error) {
	switch modelType {
	case "mongo":
		mongoClient, ok := client.(*mongo.Client)
		if !ok {
			return nil, fmt.Errorf("invalid client type for mongo: %T", client)
		}
		return NewMongoModel(mongoClient, dbName, dataSourceName), nil
	case "postgres":
		sqlDB, ok := client.(*sql.DB)
		if !ok {
			return nil, fmt.Errorf("invalid client type for postgres: %T", client)
		}
		return NewPostgresModel(sqlDB, dataSourceName), nil
	default:
		return nil, fmt.Errorf("unsupported model type: %s", modelType)
	}
}

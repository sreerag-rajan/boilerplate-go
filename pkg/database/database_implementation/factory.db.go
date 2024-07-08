package databaseimplementation

import (
	"context"
	"fmt"
	"os"
)

type DB interface {
	Connect(ctx context.Context) (interface{}, error)
}

func GetDatabase(ctx context.Context) (interface{}, error) {
	dbType := os.Getenv("DB_TYPE")
	dbUri := os.Getenv("DB_URI")
	var db DB
	if dbType == "mongo" {
		db = NewMongoDB(dbUri)
	} else if dbType == "postgres" {
		db = NewPostgresDB(dbUri)
	} else {
		return nil, fmt.Errorf("unsupported DB type")
	}
	return db.Connect(ctx)
}

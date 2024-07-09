package config

import (
	"context"

	databaseimplementation "github.com/sreerag-rajan/boilerplate-go/pkg/database/database_implementation"
)

var Client interface{}

func DBInit() error {
	ctx := context.Background()
	db, err := databaseimplementation.GetDatabase(ctx)

	Client = db

	if err != nil {
		return err
	}

	println("Connected to database!")
	return nil
}

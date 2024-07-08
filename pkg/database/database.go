package database

import (
	"context"
)

type DB interface {
	Connect(ctx context.Context) (interface{}, error)
}

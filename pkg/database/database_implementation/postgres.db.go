package databaseimplementation

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DSN string
}

func (p *PostgresDB) Connect(ctx context.Context) (interface{}, error) {
	Client, err := sql.Open("postgres", p.DSN)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgreSQL!")
	return Client, nil
}

func NewPostgresDB(connectionString string) *PostgresDB {
	return &PostgresDB{
		DSN: connectionString,
	}
}

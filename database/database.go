package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXDB struct {
	pool *pgxpool.Pool
}

func NewDatabase(connString string) (*PGXDB, error) {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config: %w", err)
	}

	dbConfig.MaxConns = 10
	dbConfig.MinConns = 2

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create db pool: %w", err)
	}

	return &PGXDB{pool: pool}, nil
}

func (db *PGXDB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *PGXDB) Close() error {
	db.pool.Close()

	return nil
}

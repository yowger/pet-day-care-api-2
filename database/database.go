package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yowger/pet-day-care-api-2/config"
)

// add logger ?
func New(cfg *config.Config) *pgxpool.Pool {
	dbConfig, err := pgxpool.ParseConfig(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("failed to parse db config: %v", err)
	}

	dbConfig.MaxConns = 10
	dbConfig.MinConns = 2

	pgxPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	// apply re tries
	if err != nil {
		log.Fatalf("failed to create db pool: %v", err)
	}

	return pgxPool
}

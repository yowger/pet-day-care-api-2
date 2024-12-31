package seeders

import (
	"context"

	"github.com/yowger/pet-day-care-api-2/config"
	"github.com/yowger/pet-day-care-api-2/database"
	sqlc "github.com/yowger/pet-day-care-api-2/database/sqlc"
)

func seedUsers(ctx context.Context) error {
	cfg := config.LoadAppConfig()
	db := database.New(cfg)
	queries := sqlc.New(db)

	return nil
}

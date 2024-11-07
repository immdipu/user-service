package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDB(connectionString string) error {
	var err error
	dbPool, err = pgxpool.New(context.Background(), connectionString)

	if err != nil {
		return err
	}

	return nil

}

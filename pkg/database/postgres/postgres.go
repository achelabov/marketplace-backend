package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const timeout = 10 * time.Second

func NewPool(uri string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, uri)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(ctx); err != nil {
		return nil, err
	}

	return dbpool, nil
}

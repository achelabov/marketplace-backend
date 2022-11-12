package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const timeout = 10 * time.Second

func NewClient(address, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type PGrepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(connStr string) (*PGrepo, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return &PGrepo{mu: sync.Mutex{}, pool: pool}, nil
}

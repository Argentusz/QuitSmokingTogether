package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(conn string) (*Storage, error) {
	p, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}
	return &Storage{pool: p}, nil
}

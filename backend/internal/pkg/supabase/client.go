package supabase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type Client struct {
// 	db	*pgxpool.Pool
// }

func NewClient(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)

	if err != nil {
		return nil, fmt.Errorf("[SUPABASE] failed to parse database URL: %w", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("[SUPABASE] failed to create pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("[SUPABASE] failed to ping pool: %w", err)
	}

	return pool, nil
}
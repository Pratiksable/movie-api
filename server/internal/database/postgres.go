package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return pool, nil
}

func CreateTables(ctx context.Context, pool *pgxpool.Pool) error {
	query := `CREATE TABLE IF NOT EXISTS movies (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		isbn VARCHAR(50) NOT NULL UNIQUE,
		title VARCHAR(100) NOT NULL ,
		director_first_name VARCHAR(100) NOT NULL, 
		director_last_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("create movies table: %w", err)
	}

	return nil
}

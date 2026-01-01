package db

import (
	"context"
	"github.com/filipcvejic/surveyly/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DB struct {
	Pool  *pgxpool.Pool
	Query *sqlc.Queries
}

func NewDatabase(connStr string) *DB {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("database connection failed:", err)
	}

	query := sqlc.New(pool)

	database := DB{
		Pool:  pool,
		Query: query,
	}

	return &database
}

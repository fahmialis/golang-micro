package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Postgres connection error:", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Postgres ping error:", err)
	}

	log.Println("PostgreSQL connected")
	return pool
}

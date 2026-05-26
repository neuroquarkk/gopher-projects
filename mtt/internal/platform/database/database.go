package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnDB(ctx context.Context, connString string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Error: unable to connect to database %v\n", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Error: failed to ping database %v\n", err)
	}

	createTable(ctx, pool)

	log.Println("Database connected successfully")
	return pool
}

func createTable(ctx context.Context, conn *pgxpool.Pool) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	task TEXT NOT NULL
	)
	`

	if _, err := conn.Exec(ctx, query); err != nil {
		log.Fatalf("Error: failed to create database table %v\n", err)
	}
}

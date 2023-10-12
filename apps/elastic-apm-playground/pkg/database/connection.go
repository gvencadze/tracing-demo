package database

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"go.elastic.co/apm/module/apmpgxv5/v2"
	"log"
)

func Connect() (*pgxpool.Pool, *sql.DB) {
	config, err := pgxpool.ParseConfig("postgres://postgres:hunter2@localhost:5432/test_db")
	if err != nil {
		log.Panic(err)
		return nil, nil
	}

	apmpgxv5.Instrument(config.ConnConfig)

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err.Error())
		return nil, nil
	}

	return pool, stdlib.OpenDB(*config.ConnConfig)
}

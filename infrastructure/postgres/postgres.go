package postgres

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var (
	pgpool *pgxpool.Pool
	pgOnce sync.Once
)

func ConnectDB() (*pgxpool.Pool, error) {
	pgOnce.Do(func() {
		var connString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_NAME"),
		)

		var err error
		config, err := pgxpool.ParseConfig(connString)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to parse database connection string")
			os.Exit(1)
		}

		pgpool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to connect to database")
			os.Exit(1)
		}
	})

	return pgpool, nil
}

func CloseDB() {
	if pgpool != nil {
		pgpool.Close()
	}
}

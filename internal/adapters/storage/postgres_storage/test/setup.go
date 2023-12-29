package test

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/multierr"
	"log/slog"
	"time"
)

const (
	DbUser     = "postgres"
	DbPassword = "12345"
	DbName     = "test_db"
)

func SetupTestDatabase() (testcontainers.Container, *sqlx.DB, error) {
	slog.Info("start DB container...")
	dbContainer, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("docker.io/postgres:15.4"),
		postgres.WithDatabase(DbName),
		postgres.WithUsername(DbUser),
		postgres.WithPassword(DbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, nil, err
	}
	slog.Info("DB container started")

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		return nil, nil, multierr.Combine(err, dbContainer.Terminate(context.Background()))
	}
	slog.Info(fmt.Sprintf("DB container host: %s", host))

	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, multierr.Combine(err, dbContainer.Terminate(context.Background()))
	}
	slog.Info(fmt.Sprintf("DB container port: %d", port.Int()))

	slog.Info("connecting to database...")
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port.Int(), DbUser, DbPassword, DbName,
	)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, nil, multierr.Combine(err, dbContainer.Terminate(context.Background()))
	}
	slog.Info("connecting to database complete")

	return dbContainer, db, nil
}

func TeardownTestDatabase(dbContainer testcontainers.Container, db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return dbContainer.Terminate(context.Background())
}

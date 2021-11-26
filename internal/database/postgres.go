package database

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/pkg/errors"
)

const newPostgresLogTag = "NewPostgres"

// StatementBuilder is a placeholder for queries
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

var errorFailedCreateConnection = errors.New("create connection failed after all attempting")

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn, driver string, attempts int) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: sqlx.Open failed to create database connection", newPostgresLogTag),
			"error", err,
		)

		return nil, err
	}

	for i := 0; i < attempts; i++ {
		err = db.PingContext(ctx)
		if err == nil {
			return db, nil
		}

		logger.ErrorKV(ctx, fmt.Sprintf("%s: b.PingContext attempt failed", newPostgresLogTag),
			"error", err,
		)
	}

	logger.ErrorKV(ctx, fmt.Sprintf("%s: create database connection failed", newPostgresLogTag),
		"error", errorFailedCreateConnection,
	)

	return nil, errorFailedCreateConnection
}

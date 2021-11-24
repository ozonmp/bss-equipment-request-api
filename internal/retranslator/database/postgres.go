package database

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/logger"
)

const newPostgresLogTag = "NewPostgres"

// StatementBuilder is a placeholder for queries
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: sqlx.Open failed to create database connection", newPostgresLogTag),
			"error", err,
		)

		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: db.PingContext failed ping the database", newPostgresLogTag),
			"error", err,
		)

		return nil, err
	}

	return db, nil
}

package database

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//WithTxFuncReturnUint64 is a function that should be run in transaction and returns uint64
type WithTxFuncReturnUint64 func(ctx context.Context, tx *sqlx.Tx) (uint64, error)

//WithTxFuncReturnBool is a function that should be run in transaction and returns bool
type WithTxFuncReturnBool func(ctx context.Context, tx *sqlx.Tx) (bool, error)

//WithTxReturnEventsFunc is a function that should be run in transaction and returns array of model.EquipmentRequestEvent
type WithTxReturnEventsFunc func(ctx context.Context, tx *sqlx.Tx) ([]model.EquipmentRequestEvent, error)

//WithTxReturnUint64 transaction for WithTxFuncReturnUint64
func WithTxReturnUint64(ctx context.Context, db *sqlx.DB, fn WithTxFuncReturnUint64) (uint64, error) {
	t, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, errors.Wrap(err, "db.BeginTxx()")
	}

	var result uint64
	if result, err = fn(ctx, t); err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return 0, errors.Wrap(err, "Tx.Rollback")
		}
		return 0, errors.Wrap(err, "Tx.WithTxFunc")
	}

	if err = t.Commit(); err != nil {
		return 0, errors.Wrap(err, "Tx.Commit")
	}

	return result, nil
}

//WithTxReturnBool transaction for WithTxFuncReturnBool
func WithTxReturnBool(ctx context.Context, db *sqlx.DB, fn WithTxFuncReturnBool) (bool, error) {
	t, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return false, errors.Wrap(err, "db.BeginTxx()")
	}

	var result bool
	if result, err = fn(ctx, t); err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return false, errors.Wrap(err, "Tx.Rollback")
		}
		return false, errors.Wrap(err, "Tx.WithTxFunc")
	}

	if err = t.Commit(); err != nil {
		return false, errors.Wrap(err, "Tx.Commit")
	}

	return result, nil
}

//WithTxReturnEvents transaction for WithTxReturnEventsFunc
func WithTxReturnEvents(ctx context.Context, db *sqlx.DB, fn WithTxReturnEventsFunc) ([]model.EquipmentRequestEvent, error) {
	t, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "db.BeginTxx()")
	}

	var result []model.EquipmentRequestEvent
	if result, err = fn(ctx, t); err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return nil, errors.Wrap(err, "Tx.Rollback")
		}
		return nil, errors.Wrap(err, "Tx.WithTxFunc")
	}

	if err = t.Commit(); err != nil {
		return nil, errors.Wrap(err, "Tx.Commit")
	}
	return result, nil
}

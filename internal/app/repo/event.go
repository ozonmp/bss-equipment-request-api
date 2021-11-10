package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

// EventRepo is DAO for Equipment Request Events
type EventRepo interface {
	Lock(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error
	Remove(ctx context.Context, eventIDs []uint64) error
}

type eventRepo struct {
	db *sqlx.DB
}

// NewEventRepo returns Repo interface
func NewEventRepo(db *sqlx.DB) EventRepo {
	return &eventRepo{
		db: db,
	}
}

func (r *eventRepo) Lock(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
	events, err := database.WithTxReturnEvents(ctx, db, func(ctx context.Context, tx *sqlx.Tx) ([]model.EquipmentRequestEvent, error) {
		subSel := database.StatementBuilder.
			Select("id").
			From("equipment_request_event").
			Where(sq.Eq{"status": model.Unlocked}).
			Limit(batchSize).
			OrderBy("id").
			Suffix("FOR UPDATE SKIP LOCKED")

		sb := database.StatementBuilder.
			Update("equipment_request_event").
			Where(subSel.Prefix("id IN (").Suffix(")")).
			Set("status", model.Locked).
			Set("updated_at", time.Now()).
			Suffix("RETURNING equipment_request_events.*")

		query, args, err := sb.ToSql()
		if err != nil {
			return nil, err
		}

		var queryer sqlx.QueryerContext
		if tx == nil {
			queryer = r.db
		} else {
			queryer = tx
		}

		rows, err := queryer.QueryContext(ctx, query, args...)
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Error().Err(err).Msg("Lock - rows.Close()")
			}
		}(rows)

		if err != nil {
			return nil, errors.Wrap(err, "db.QueryContext()")
		}

		equipmentRequestEvents := make([]model.EquipmentRequestEvent, 0)
		err = sqlx.StructScan(rows, &equipmentRequestEvents)

		if err != nil {
			return nil, err
		}

		return equipmentRequestEvents, nil
	})

	return events, err
}

func (r *eventRepo) Remove(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.
		Update("equipment_request_event").
		Where(sq.And{
			sq.Eq{"id": eventIDs},
			sq.Eq{"deleted_at": nil}}).
		Set("status", model.Processed).
		Set("updated_at", time.Now())

	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(err, "db.ExecContext()")
	}

	return nil
}

func (r *eventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {
	sb := database.StatementBuilder.
		Update("equipment_request_event").
		Where(sq.And{
			sq.Eq{"id": eventIDs},
			sq.Eq{"deleted_at": nil},
			sq.NotEq{"status": model.Processed}}).
		Set("status", model.Unlocked).
		Set("updated_at", time.Now())

	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(err, "db.ExecContext()")
	}

	return nil
}

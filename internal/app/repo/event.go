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

const (
	equipmentRequestEventTable           = "equipment_request_event"
	equipmentRequestEventIDColumn        = "id"
	equipmentRequestEventStatusColumn    = "status"
	equipmentRequestEventUpdatedAtColumn = "updated_at"
	equipmentRequestEventDeletedAtColumn = "deleted_at"
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
			Select(equipmentRequestEventIDColumn).
			From(equipmentRequestEventTable).
			Where(sq.Eq{equipmentRequestEventStatusColumn: model.Unlocked}).
			Limit(batchSize).
			OrderBy(equipmentRequestEventIDColumn).
			Suffix("FOR UPDATE SKIP LOCKED")

		sb := database.StatementBuilder.
			Update(equipmentRequestEventTable).
			Where(subSel.Prefix(equipmentRequestEventIDColumn+" IN (").Suffix(")")).
			Set(equipmentRequestEventStatusColumn, model.Locked).
			Set(equipmentRequestEventUpdatedAtColumn, time.Now()).
			Suffix("RETURNING *")

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

		if err != nil {
			return nil, errors.Wrap(err, "db.QueryContext()")
		}

		defer func(rows *sql.Rows) {
			err = rows.Close()
			if err != nil {
				log.Error().Err(err).Msg("Lock - rows.Close()")
			}
		}(rows)

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
		Update(equipmentRequestEventTable).
		Where(sq.And{
			sq.Eq{equipmentRequestEventIDColumn: eventIDs},
			sq.Eq{equipmentRequestEventDeletedAtColumn: nil}}).
		Set(equipmentRequestEventStatusColumn, model.Processed).
		Set(equipmentRequestEventUpdatedAtColumn, time.Now())

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
		Update(equipmentRequestEventTable).
		Where(sq.And{
			sq.Eq{equipmentRequestEventIDColumn: eventIDs},
			sq.Eq{equipmentRequestEventDeletedAtColumn: nil},
			sq.NotEq{equipmentRequestEventStatusColumn: model.Processed}}).
		Set(equipmentRequestEventStatusColumn, model.Unlocked).
		Set(equipmentRequestEventUpdatedAtColumn, time.Now())

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

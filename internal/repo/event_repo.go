package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

// EventRepo is DAO for Equipment Request Events
type EventRepo interface {
	Add(ctx context.Context, event *model.EquipmentRequestEvent, tx *sqlx.Tx) error
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

func (r *eventRepo) Add(ctx context.Context, event *model.EquipmentRequestEvent, tx *sqlx.Tx) error {
	convertedPayload := event.Payload.ConvertToPb()
	payload, err := protojson.Marshal(convertedPayload)

	if err != nil {
		return err
	}

	sb := database.StatementBuilder.
		Insert("equipment_request_event").
		Columns("type", "status", "created_at", "updated_at", "equipment_request_id", "payload").
		Values(
			event.Type,
			event.Status,
			event.CreatedAt,
			event.UpdatedAt,
			event.EquipmentRequestID,
			payload,
		)

	query, args, err := sb.ToSql()

	if err != nil {
		return err
	}

	var queryer sqlx.ExecerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	_, err = queryer.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.Wrap(err, "db.ExecContext()")
	}

	return nil
}

package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	equipmentRequestEventTable           = "equipment_request_event"
	equipmentRequestEventStatusColumn    = "status"
	equipmentRequestEventCreatedAtColumn = "created_at"
	equipmentRequestEventTypeColumn      = "type"
	equipmentRequestEventRequestIDColumn = "equipment_request_id"
	equipmentRequestPayloadColumn        = "payload"
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "eventRepo.Add")
	defer span.Finish()

	sb := database.StatementBuilder.
		Insert(equipmentRequestEventTable).
		Columns(
			equipmentRequestEventTypeColumn,
			equipmentRequestEventStatusColumn,
			equipmentRequestEventCreatedAtColumn,
			equipmentRequestEventRequestIDColumn).
		Values(
			event.Type,
			event.Status,
			event.CreatedAt,
			event.EquipmentRequestID,
		)

	if event.Payload != nil {
		convertedPayload := event.Payload.ConvertToPb()
		payload, err := protojson.Marshal(convertedPayload)

		if err != nil {
			return err
		}

		sb.Columns(equipmentRequestPayloadColumn).Values(payload)
	}
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

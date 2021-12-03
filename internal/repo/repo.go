package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/pkg/errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

const (
	equipmentRequestTable             = "equipment_request"
	equipmentRequestIDColumn          = "id"
	equipmentRequestEquipmentIDColumn = "equipment_id"
	equipmentRequestEmployeeIDColumn  = "employee_id"
	equipmentRequestStatusColumn      = "equipment_request_status"
	equipmentRequestUpdatedAtColumn   = "updated_at"
	equipmentRequestCreatedAtColumn   = "created_at"
	equipmentRequestDoneAtColumn      = "done_at"
	equipmentRequestDeletedAtAtColumn = "deleted_at"
)

// EquipmentRequestRepo is DAO for Equipment Request
type EquipmentRequestRepo interface {
	CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest, tx *sqlx.Tx) (uint64, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64, tx *sqlx.Tx) (bool, error)
	Exists(ctx context.Context, equipmentRequestID uint64) (bool, error)
	UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64, tx *sqlx.Tx) (bool, error)
	UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status model.EquipmentRequestStatus, tx *sqlx.Tx) (bool, error)
}

type equipmentRequestRepo struct {
	db *sqlx.DB
}

// NewEquipmentRequestRepo returns Repo interface
func NewEquipmentRequestRepo(db *sqlx.DB) EquipmentRequestRepo {
	return &equipmentRequestRepo{
		db: db,
	}
}

func (r *equipmentRequestRepo) CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest, tx *sqlx.Tx) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.CreateEquipmentRequest")
	defer span.Finish()

	sb := database.StatementBuilder.
		Insert(equipmentRequestTable).
		Columns(
			equipmentRequestEmployeeIDColumn,
			equipmentRequestEquipmentIDColumn,
			equipmentRequestCreatedAtColumn,
			equipmentRequestUpdatedAtColumn,
			equipmentRequestDeletedAtAtColumn,
			equipmentRequestDoneAtColumn,
			equipmentRequestStatusColumn).
		Values(
			equipmentRequest.EmployeeID,
			equipmentRequest.EquipmentID,
			equipmentRequest.CreatedAt,
			equipmentRequest.UpdatedAt,
			equipmentRequest.DeletedAt,
			equipmentRequest.DoneAt,
			equipmentRequest.EquipmentRequestStatus,
		).Suffix("RETURNING " + equipmentRequestIDColumn)

	query, args, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	var id uint64
	err = queryer.QueryRowxContext(ctx, query, args...).Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, errors.Wrap(err, "db.QueryRowxContext()")
	}

	return id, nil
}

func (r *equipmentRequestRepo) RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64, tx *sqlx.Tx) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.RemoveEquipmentRequest")
	defer span.Finish()

	sb := database.StatementBuilder.
		Update(equipmentRequestTable).
		Set(equipmentRequestDeletedAtAtColumn, time.Now()).
		Where(sq.And{
			sq.Eq{equipmentRequestIDColumn: equipmentRequestID},
			sq.Eq{equipmentRequestDeletedAtAtColumn: nil}})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, err
	}

	var queryer sqlx.ExecerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	var result sql.Result
	result, err = queryer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "repo.RowsAffected()")
	}

	if affected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *equipmentRequestRepo) Exists(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Exists")
	defer span.Finish()

	sb := database.StatementBuilder.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(equipmentRequestTable).
		Where(sq.And{
			sq.Eq{equipmentRequestIDColumn: equipmentRequestID},
			sq.Eq{equipmentRequestDeletedAtAtColumn: nil}}).
		Suffix(")")

	query, args, err := sb.ToSql()
	if err != nil {
		return false, err
	}

	var exists bool
	err = r.db.QueryRowxContext(ctx, query, args...).Scan(&exists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, errors.Wrap(err, "db.QueryRowxContext()")
	}

	return exists, nil
}

// nolint:dupl
func (r *equipmentRequestRepo) UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64, tx *sqlx.Tx) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.UpdateEquipmentIDEquipmentRequest")
	defer span.Finish()

	sb := database.StatementBuilder.
		Update(equipmentRequestTable).
		Set(equipmentRequestUpdatedAtColumn, time.Now()).
		Set(equipmentRequestEquipmentIDColumn, equipmentID).
		Where(sq.And{
			sq.Eq{equipmentRequestIDColumn: equipmentRequestID},
			sq.Eq{equipmentRequestDeletedAtAtColumn: nil}})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, err
	}

	var queryer sqlx.ExecerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	var result sql.Result
	result, err = queryer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "repo.RowsAffected()")
	}

	if affected == 0 {
		return false, nil
	}

	return true, nil
}

// nolint:dupl
func (r *equipmentRequestRepo) UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status model.EquipmentRequestStatus, tx *sqlx.Tx) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.UpdateStatusEquipmentRequest")
	defer span.Finish()

	sb := database.StatementBuilder.
		Update(equipmentRequestTable).
		Set(equipmentRequestUpdatedAtColumn, time.Now()).
		Set(equipmentRequestStatusColumn, status).
		Where(sq.And{
			sq.Eq{equipmentRequestIDColumn: equipmentRequestID},
			sq.Eq{equipmentRequestDeletedAtAtColumn: nil}})

	query, args, err := sb.ToSql()
	if err != nil {
		return false, err
	}

	var queryer sqlx.ExecerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	var result sql.Result
	result, err = queryer.ExecContext(ctx, query, args...)

	if err != nil {
		return false, errors.Wrap(err, "db.SelectContext()")
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return false, errors.Wrap(err, "repo.RowsAffected()")
	}

	if affected == 0 {
		return false, nil
	}

	return true, nil
}

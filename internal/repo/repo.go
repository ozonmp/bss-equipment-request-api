package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
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
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error)
	CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest, tx *sqlx.Tx) (uint64, error)
	ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]model.EquipmentRequest, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64, tx *sqlx.Tx) (bool, error)
	Exists(ctx context.Context, equipmentRequestID uint64) (bool, error)
	UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64, tx *sqlx.Tx) (bool, error)
	UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status model.EquipmentRequestStatus, tx *sqlx.Tx) (bool, error)
}

type equipmentRequestRepo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewEquipmentRequestRepo returns Repo interface
func NewEquipmentRequestRepo(db *sqlx.DB, batchSize uint) EquipmentRequestRepo {
	return &equipmentRequestRepo{
		db:        db,
		batchSize: batchSize,
	}
}

func (r *equipmentRequestRepo) DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error) {
	sb := database.StatementBuilder.
		Select("*").
		From(equipmentRequestTable).
		Where(sq.And{
			sq.Eq{equipmentRequestIDColumn: equipmentRequestID},
			sq.Eq{equipmentRequestDeletedAtAtColumn: nil}})

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	var equipmentRequest model.EquipmentRequest
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&equipmentRequest)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "db.QueryRowxContext()")
	}

	return &equipmentRequest, nil
}

func (r *equipmentRequestRepo) CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest, tx *sqlx.Tx) (uint64, error) {
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
		return 0, errors.Wrap(err, "db.QueryRowContext()")
	}

	return id, nil
}

func (r *equipmentRequestRepo) ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]model.EquipmentRequest, error) {
	sb := database.StatementBuilder.
		Select("*").
		From(equipmentRequestTable).
		Where(sq.Eq{equipmentRequestDeletedAtAtColumn: nil}).
		OrderBy(equipmentRequestIDColumn).
		Limit(limit).
		Offset(offset)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	var equipmentRequests []model.EquipmentRequest
	err = r.db.SelectContext(ctx, &equipmentRequests, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return equipmentRequests, nil
}

func (r *equipmentRequestRepo) RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64, tx *sqlx.Tx) (bool, error) {
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

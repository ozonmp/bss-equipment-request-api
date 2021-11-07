package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	internal_errors "github.com/ozonmp/bss-equipment-request-api/internal/pkg/errors"

	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

// Repo is DAO for Equipment Request
type Repo interface {
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error)
	CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest) (uint64, error)
	ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error)
	Exists(ctx context.Context, equipmentRequestID uint64) (bool, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error) {
	return nil, internal_errors.ErrNotImplementedMethod
}

func (r *repo) CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest) (uint64, error) {
	return 0, internal_errors.ErrNotImplementedMethod
}

func (r *repo) ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error) {
	return nil, internal_errors.ErrNotImplementedMethod
}

func (r *repo) RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	return false, internal_errors.ErrNotImplementedMethod
}

func (r *repo) Exists(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	return false, internal_errors.ErrNotImplementedMethod
}

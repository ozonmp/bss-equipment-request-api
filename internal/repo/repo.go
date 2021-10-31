package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

// Repo is DAO for Equipment Request
type Repo interface {
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestId uint64) (*model.EquipmentRequest, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeEquipmentRequest(ctx context.Context, equipmentRequestId uint64) (*model.EquipmentRequest, error) {
	return nil, nil
}

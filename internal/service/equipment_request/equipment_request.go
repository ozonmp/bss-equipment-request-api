package equipment_request

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	internal_errors "github.com/ozonmp/bss-equipment-request-api/internal/pkg/errors"
	"github.com/ozonmp/bss-equipment-request-api/internal/repo"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type service struct {
	repository repo.Repo
}

// ServiceInterface is a interface for equipment request service
type ServiceInterface interface {
	DescribeEquipmentRequest(ctx context.Context, id uint64) (*model.EquipmentRequest, error)
	CreateEquipmentRequest(
		ctx context.Context,
		equipmentID uint64,
		employeeID uint64,
		createdAt *timestamppb.Timestamp,
		doneAt *timestamppb.Timestamp,
		equipmentRequestStatusID pb.EquipmentRequestStatus) (uint64, error)
	ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error)
	RemoveEquipmentRequest(ctx context.Context, id uint64) (bool, error)
	CheckExistsEquipmentRequest(ctx context.Context, id uint64) (bool, error)
}

// New is a function to create a new service
func New(repository repo.Repo) ServiceInterface {
	return service{
		repository: repository,
	}
}

// ErrNoEquipmentRequest is a "equipment request not founded" error
var ErrNoEquipmentRequest = errors.Wrap(internal_errors.ErrNotFound, "no equipment request")

// ErrNoCreatedEquipmentRequest is a "unable to create equipment request" error
var ErrNoCreatedEquipmentRequest = errors.Wrap(internal_errors.ErrNotCreated, "unable to create equipment request")

// ErrNoListEquipmentRequest is a "unable to get list of equipment requests" error
var ErrNoListEquipmentRequest = errors.Wrap(internal_errors.ErrNotFound, "unable to get list of equipment requests")

// ErrNoRemovedEquipmentRequest is a "unable to remove equipment request" error
var ErrNoRemovedEquipmentRequest = errors.Wrap(internal_errors.ErrNotRemoved, "unable to remove equipment request")

func (s service) DescribeEquipmentRequest(ctx context.Context, id uint64) (*model.EquipmentRequest, error) {
	equipmentRequest, err := s.repository.DescribeEquipmentRequest(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.DescribeEquipmentRequest")
	}

	if equipmentRequest == nil {
		return nil, ErrNoEquipmentRequest
	}

	return equipmentRequest, nil
}

func (s service) CreateEquipmentRequest(
	ctx context.Context,
	equipmentID uint64,
	employeeID uint64,
	createdAt *timestamppb.Timestamp,
	doneAt *timestamppb.Timestamp,
	equipmentRequestStatusID pb.EquipmentRequestStatus) (uint64, error) {

	var newCreatedAt time.Time
	var newDoneAt time.Time

	if createdAt != nil {
		newCreatedAt = createdAt.AsTime()
	}

	if doneAt != nil {
		newDoneAt = doneAt.AsTime()
	}

	newEquipmentRequest := &model.EquipmentRequest{
		EquipmentID:              equipmentID,
		EmployeeID:               employeeID,
		CreatedAt:                &newCreatedAt,
		DoneAt:                   &newDoneAt,
		EquipmentRequestStatusID: model.EquipmentRequestStatus(equipmentRequestStatusID),
	}

	id, err := s.repository.CreateEquipmentRequest(ctx, newEquipmentRequest)
	if err != nil {
		return 0, errors.Wrap(err, "repository.CreateEquipmentRequest")
	}

	if id == 0 {
		return 0, ErrNoCreatedEquipmentRequest
	}

	return id, nil
}

func (s service) ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error) {

	equipmentRequests, err := s.repository.ListEquipmentRequest(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.ListEquipmentRequest")
	}

	if equipmentRequests == nil {
		return nil, ErrNoListEquipmentRequest
	}

	return equipmentRequests, nil
}

func (s service) CheckExistsEquipmentRequest(ctx context.Context, id uint64) (bool, error) {

	exists, err := s.repository.Exists(ctx, id)

	if err != nil {
		return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
	}

	if exists == false {
		return false, ErrNoEquipmentRequest
	}

	return exists, nil
}

func (s service) RemoveEquipmentRequest(ctx context.Context, id uint64) (bool, error) {

	result, err := s.repository.RemoveEquipmentRequest(ctx, id)

	if err != nil {
		return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
	}

	if result == false {
		return false, ErrNoRemovedEquipmentRequest
	}

	return result, nil
}

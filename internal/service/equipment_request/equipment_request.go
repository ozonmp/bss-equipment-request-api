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

type Service struct {
	repository repo.Repo
}

type ServiceInterface interface {
	DescribeEquipmentRequest(ctx context.Context, id uint64) (*model.EquipmentRequest, error)
	CreateEquipmentRequest(
		ctx context.Context,
		equipmentId uint64,
		employeeId uint64,
		createdAt *timestamppb.Timestamp,
		doneAt *timestamppb.Timestamp,
		equipmentRequestStatusId pb.EquipmentRequestStatus) (uint64, error)
	ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error)
	RemoveEquipmentRequest(ctx context.Context, id uint64) (bool, error)
	CheckExistsEquipmentRequest(ctx context.Context, id uint64) (bool, error)
}

func New(repository repo.Repo) Service {
	return Service{
		repository: repository,
	}
}

var ErrNoEquipmentRequest = errors.Wrap(internal_errors.ErrNotFound, "no equipment request")
var ErrNoCreatedEquipmentRequest = errors.Wrap(internal_errors.ErrNotCreated, "unable to create equipment request")
var ErrNoListEquipmentRequest = errors.Wrap(internal_errors.ErrNotFound, "unable to get list of equipment requests")
var ErrNoRemovedEquipmentRequest = errors.Wrap(internal_errors.ErrNotRemoved, "unable to remove equipment request")

func (s Service) DescribeEquipmentRequest(ctx context.Context, id uint64) (*model.EquipmentRequest, error) {
	equipmentRequest, err := s.repository.DescribeEquipmentRequest(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "repository.DescribeEquipmentRequest")
	}

	if equipmentRequest == nil {
		return nil, ErrNoEquipmentRequest
	}

	return equipmentRequest, nil
}

func (s Service) CreateEquipmentRequest(
	ctx context.Context,
	equipmentId uint64,
	employeeId uint64,
	createdAt *timestamppb.Timestamp,
	doneAt *timestamppb.Timestamp,
	equipmentRequestStatusId pb.EquipmentRequestStatus) (uint64, error) {

	var newCreatedAt time.Time
	var newDoneAt time.Time

	if createdAt != nil {
		newCreatedAt = createdAt.AsTime()
	}

	if doneAt != nil {
		newDoneAt = doneAt.AsTime()
	}

	newEquipmentRequest := &model.EquipmentRequest{
		EquipmentID:              equipmentId,
		EmployeeID:               employeeId,
		CreatedAt:                &newCreatedAt,
		DoneAt:                   &newDoneAt,
		EquipmentRequestStatusID: model.EquipmentRequestStatus(equipmentRequestStatusId),
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

func (s Service) ListEquipmentRequest(ctx context.Context) ([]model.EquipmentRequest, error) {

	equipmentRequests, err := s.repository.ListEquipmentRequest(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.ListEquipmentRequest")
	}

	if equipmentRequests == nil {
		return nil, ErrNoListEquipmentRequest
	}

	return equipmentRequests, nil
}

func (s Service) CheckExistsEquipmentRequest(ctx context.Context, id uint64) (bool, error) {

	exists, err := s.repository.Exists(ctx, id)

	if err != nil {
		return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
	}

	if exists == false {
		return false, ErrNoEquipmentRequest
	}

	return exists, nil
}

func (s Service) RemoveEquipmentRequest(ctx context.Context, id uint64) (bool, error) {

	result, err := s.repository.RemoveEquipmentRequest(ctx, id)

	if err != nil {
		return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
	}

	if result == false {
		return false, ErrNoRemovedEquipmentRequest
	}

	return result, nil
}

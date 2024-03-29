package equipment_request

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/metrics"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/repo"
	"github.com/pkg/errors"
)

type service struct {
	db                *sqlx.DB
	requestRepository repo.EquipmentRequestRepo
	eventRepository   repo.EventRepo
}

// ServiceInterface is a interface for equipment request service
type ServiceInterface interface {
	DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error)
	CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest) (uint64, error)
	ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]model.EquipmentRequest, error)
	RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error)
	CheckExistsEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error)
	UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error)
	UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status model.EquipmentRequestStatus) (bool, error)
}

// New is a function to create a new service
func New(db *sqlx.DB, requestRepository repo.EquipmentRequestRepo, eventRepository repo.EventRepo) ServiceInterface {
	return service{
		db:                db,
		requestRepository: requestRepository,
		eventRepository:   eventRepository,
	}
}

// ErrNoExistsEquipmentRequest is a "equipment request not founded" error
var ErrNoExistsEquipmentRequest = errors.New("equipment request with this id does not exist")

// ErrNoCreatedEquipmentRequest is a "unable to create equipment request" error
var ErrNoCreatedEquipmentRequest = errors.New("unable to create equipment request")

// ErrNoListEquipmentRequest is a "unable to get list of equipment requests" error
var ErrNoListEquipmentRequest = errors.New("unable to get list of equipment requests")

// ErrNoRemovedEquipmentRequest is a "unable to remove equipment request" error
var ErrNoRemovedEquipmentRequest = errors.New("unable to remove equipment request")

// ErrNoUpdatedEquipmentIDEquipmentRequest is a "unable to update equipment id of equipment request" error
var ErrNoUpdatedEquipmentIDEquipmentRequest = errors.New("unable to update equipment id of equipment request")

// ErrNoUpdatedStatusEquipmentRequest is a "unable to update status of equipment request" error
var ErrNoUpdatedStatusEquipmentRequest = errors.New("unable to update status of equipment request")

func (s service) DescribeEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (*model.EquipmentRequest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.DescribeEquipmentRequest")
	defer span.Finish()

	equipmentRequest, err := s.requestRepository.DescribeEquipmentRequest(ctx, equipmentRequestID)
	if err != nil {
		return nil, errors.Wrap(err, "repository.DescribeEquipmentRequest")
	}

	return equipmentRequest, nil
}

func (s service) CreateEquipmentRequest(ctx context.Context, equipmentRequest *model.EquipmentRequest) (uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateEquipmentRequest")
	defer span.Finish()

	createdRequestID, txErr := database.WithTxReturnUint64(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) (uint64, error) {
		id, err := s.requestRepository.CreateEquipmentRequest(ctx, equipmentRequest, tx)
		if err != nil {
			return 0, errors.Wrap(err, "requestRepository.CreateEquipmentRequest")
		}

		if id == 0 {
			return 0, ErrNoCreatedEquipmentRequest
		}
		equipmentRequest.ID = id

		event, err := model.FormCreatedEvent(equipmentRequest)

		if err != nil {
			return 0, errors.Wrap(err, "model.FormCreatedEvent")
		}

		err = s.eventRepository.Add(ctx, event, tx)

		if err != nil {
			return 0, errors.Wrap(err, "eventRepository.Add")
		}

		metrics.IncTotalCudOperations(event.Type)

		return id, nil
	})

	if txErr != nil {
		return createdRequestID, txErr
	}

	return createdRequestID, nil
}

func (s service) ListEquipmentRequest(ctx context.Context, limit uint64, offset uint64) ([]model.EquipmentRequest, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.ListEquipmentRequest")
	defer span.Finish()

	equipmentRequests, err := s.requestRepository.ListEquipmentRequest(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "repository.ListEquipmentRequest")
	}

	if equipmentRequests == nil {
		return nil, ErrNoListEquipmentRequest
	}

	return equipmentRequests, nil
}

func (s service) CheckExistsEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CheckExistsEquipmentRequest")
	defer span.Finish()

	exists, err := s.requestRepository.Exists(ctx, equipmentRequestID)

	if err != nil {
		return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
	}

	if !exists {
		return false, ErrNoExistsEquipmentRequest
	}

	return exists, nil
}

func (s service) RemoveEquipmentRequest(ctx context.Context, equipmentRequestID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.RemoveEquipmentRequest")
	defer span.Finish()

	deleted, txErr := database.WithTxReturnBool(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) (bool, error) {
		result, err := s.requestRepository.RemoveEquipmentRequest(ctx, equipmentRequestID, tx)
		if err != nil {
			return false, errors.Wrap(err, "repository.RemoveEquipmentRequest")
		}

		if !result {
			return false, ErrNoRemovedEquipmentRequest
		}

		event := model.FormRemovedEvent(equipmentRequestID)

		err = s.eventRepository.Add(ctx, event, tx)

		if err != nil {
			return false, errors.Wrap(err, "eventRepository.Add")
		}

		metrics.IncTotalCudOperations(event.Type)

		return result, nil
	})

	if txErr != nil {
		return deleted, txErr
	}

	return deleted, nil
}

func (s service) UpdateEquipmentIDEquipmentRequest(ctx context.Context, equipmentRequestID uint64, equipmentID uint64) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UpdateEquipmentIDEquipmentRequest")
	defer span.Finish()

	updated, txErr := database.WithTxReturnBool(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) (bool, error) {
		result, err := s.requestRepository.UpdateEquipmentIDEquipmentRequest(ctx, equipmentRequestID, equipmentID, tx)
		if err != nil {
			return false, errors.Wrap(err, "repository.UpdateEquipmentIDEquipmentRequest")
		}

		if !result {
			return false, ErrNoUpdatedEquipmentIDEquipmentRequest
		}

		event := model.FormUpdatedEquipmentIDEvent(equipmentRequestID, equipmentID)

		err = s.eventRepository.Add(ctx, event, tx)

		if err != nil {
			return false, errors.Wrap(err, "eventRepository.Add")
		}

		metrics.IncTotalCudOperations(event.Type)

		return result, nil
	})

	if txErr != nil {
		return updated, txErr
	}

	return updated, nil
}

func (s service) UpdateStatusEquipmentRequest(ctx context.Context, equipmentRequestID uint64, status model.EquipmentRequestStatus) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UpdateStatusEquipmentRequest")
	defer span.Finish()

	updated, txErr := database.WithTxReturnBool(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) (bool, error) {
		result, err := s.requestRepository.UpdateStatusEquipmentRequest(ctx, equipmentRequestID, status, tx)
		if err != nil {
			return false, errors.Wrap(err, "repository.UpdateStatusEquipmentRequest")
		}

		if !result {
			return false, ErrNoUpdatedStatusEquipmentRequest
		}

		event, err := model.FormUpdatedStatusEvent(equipmentRequestID, status)

		if err != nil {
			return false, errors.Wrap(err, "model.FormUpdatedStatusEvent")
		}

		err = s.eventRepository.Add(ctx, event, tx)

		if err != nil {
			return false, errors.Wrap(err, "eventRepository.Add")
		}

		metrics.IncTotalCudOperations(event.Type)

		return result, nil
	})

	if txErr != nil {
		return updated, txErr
	}

	return updated, nil
}

package model

import (
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// EventType is a type of events
type EventType string

const (
	// Created is a "created item" type of events
	Created EventType = "EQUIPMENT_REQUEST_EVENT_TYPE_CREATED"
	// UpdatedEquipmentId is a "updated equipment id of item" type of events
	UpdatedEquipmentId = "EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_EQUIPMENT_ID"
	// UpdatedStatus is a "updated status of item" type of events
	UpdatedStatus = "EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_STATUS"
	// Removed is a "removed item" type of events
	Removed = "EQUIPMENT_REQUEST_EVENT_TYPE_DELETED"
)

// EventStatus is a status of events
type EventStatus string

const (
	// Locked is a locked status of events
	Locked EventStatus = "EQUIPMENT_REQUEST_EVENT_STATUS_LOCKED"
	// Unlocked is a unlocked status of events
	Unlocked EventStatus = "EQUIPMENT_REQUEST_EVENT_STATUS_UNLOCKED"
	// Processed is a precessed status of events
	Processed = "EQUIPMENT_REQUEST_EVENT_STATUS_PROCESSED"
)

// EquipmentRequestEvent is a event of equipment request
type EquipmentRequestEvent struct {
	ID                 uint64                        `db:"id"`
	Type               EventType                     `db:"type"`
	Status             EventStatus                   `db:"status"`
	CreatedAt          time.Time                     `db:"created_at"`
	UpdatedAt          time.Time                     `db:"updated_at"`
	EquipmentRequestID uint64                        `db:"equipment_request_id"`
	Payload            *EquipmentRequestEventPayload `db:"payload"`
}

type EquipmentRequestEventPayload pb.EquipmentRequest

func (e *EquipmentRequestEventPayload) Scan(src interface{}) (err error) {
	var eqp pb.EquipmentRequest

	switch src.(type) {
	case string:
		err = protojson.Unmarshal([]byte(src.(string)), &eqp)
	case []byte:
		err = protojson.Unmarshal(src.([]byte), &eqp)
	default:
		return errors.New("Incompatible type for EquipmentRequestEventPayload")
	}

	if err != nil {
		return err
	}

	*e = EquipmentRequestEventPayload{
		Id:                     eqp.Id,
		EquipmentId:            eqp.EquipmentId,
		EmployeeId:             eqp.EmployeeId,
		CreatedAt:              eqp.CreatedAt,
		UpdatedAt:              eqp.UpdatedAt,
		DoneAt:                 eqp.DoneAt,
		DeletedAt:              eqp.DeletedAt,
		EquipmentRequestStatus: eqp.EquipmentRequestStatus,
	}

	return nil
}

func (e *EquipmentRequestEventPayload) ConvertToPb() *pb.EquipmentRequest {
	return &pb.EquipmentRequest{
		Id:                     e.Id,
		EquipmentId:            e.EquipmentId,
		EmployeeId:             e.EmployeeId,
		CreatedAt:              e.CreatedAt,
		UpdatedAt:              e.UpdatedAt,
		DoneAt:                 e.DoneAt,
		DeletedAt:              e.DeletedAt,
		EquipmentRequestStatus: e.EquipmentRequestStatus,
	}
}

func FormCreatedEvent(request *EquipmentRequest) (*EquipmentRequestEvent, error) {
	status, ok := pb.EquipmentRequestStatus_value[string(request.EquipmentRequestStatus)]
	var equipmentRequestStatus pb.EquipmentRequestStatus

	if ok {
		equipmentRequestStatus = pb.EquipmentRequestStatus(status)
	} else {
		return nil, errors.New("unable to convert equipment request status")
	}

	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               Created,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		EquipmentRequestID: request.ID,
		Payload: &EquipmentRequestEventPayload{
			Id:                     request.ID,
			EmployeeId:             request.EmployeeID,
			EquipmentId:            request.EquipmentID,
			CreatedAt:              timestamppb.New(request.CreatedAt),
			UpdatedAt:              timestamppb.New(request.UpdatedAt),
			DoneAt:                 timestamppb.New(request.DoneAt.Time),
			DeletedAt:              timestamppb.New(request.DeletedAt.Time),
			EquipmentRequestStatus: equipmentRequestStatus,
		},
	}, nil
}

func FormUpdatedEquipmentIdEvent(requestID uint64, equipmentID uint64) *EquipmentRequestEvent {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               UpdatedEquipmentId,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		EquipmentRequestID: requestID,
		Payload: &EquipmentRequestEventPayload{
			Id:          requestID,
			EquipmentId: equipmentID,
		},
	}
}

func FormUpdatedStatusEvent(requestID uint64, status EquipmentRequestStatus) (*EquipmentRequestEvent, error) {
	statusVal, ok := pb.EquipmentRequestStatus_value[string(status)]
	var equipmentRequestStatus pb.EquipmentRequestStatus

	if ok {
		equipmentRequestStatus = pb.EquipmentRequestStatus(statusVal)
	} else {
		return nil, errors.New("unable to convert equipment request status")
	}

	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               UpdatedStatus,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		EquipmentRequestID: requestID,
		Payload: &EquipmentRequestEventPayload{
			Id:                     requestID,
			EquipmentRequestStatus: equipmentRequestStatus,
		},
	}, nil
}

func FormRemovedEvent(equipmentRequestID uint64) *EquipmentRequestEvent {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               Removed,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		EquipmentRequestID: equipmentRequestID,
	}
}

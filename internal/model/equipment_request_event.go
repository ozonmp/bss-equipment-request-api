package model

import (
	"database/sql"
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// EventType is a type of events
type EventType string

func (et EventType) String() string {
	return string(et)
}

const (
	// Created is a "created item" type of events
	Created EventType = "EQUIPMENT_REQUEST_EVENT_TYPE_CREATED"
	// UpdatedEquipmentID is a "updated equipment id of item" type of events
	UpdatedEquipmentID = "EQUIPMENT_REQUEST_EVENT_TYPE_UPDATED_EQUIPMENT_ID"
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
	UpdatedAt          sql.NullTime                  `db:"updated_at"`
	EquipmentRequestID uint64                        `db:"equipment_request_id"`
	Payload            *EquipmentRequestEventPayload `db:"payload"`
}

// EquipmentRequestEventPayload is a detailed info about event
type EquipmentRequestEventPayload pb.EquipmentRequest

// Scan EquipmentRequestEventPayload
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

// ConvertToPb is a function to convert pb.EquipmentRequest to EquipmentRequestEventPayload
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

// FormCreatedEvent is a function to create event about equipment request creating
func FormCreatedEvent(request *EquipmentRequest) (*EquipmentRequestEvent, error) {
	status, ok := pb.EquipmentRequestStatus_value[string(request.EquipmentRequestStatus)]
	var equipmentRequestStatus pb.EquipmentRequestStatus

	if !ok {
		return nil, errors.New("unable to convert equipment request status")
	}

	equipmentRequestStatus = pb.EquipmentRequestStatus(status)

	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               Created,
		CreatedAt:          time.Now(),
		EquipmentRequestID: request.ID,
		Payload: &EquipmentRequestEventPayload{
			Id:                     request.ID,
			EmployeeId:             request.EmployeeID,
			EquipmentId:            request.EquipmentID,
			CreatedAt:              timestamppb.New(request.CreatedAt),
			UpdatedAt:              timestamppb.New(request.UpdatedAt.Time),
			DoneAt:                 timestamppb.New(request.DoneAt.Time),
			DeletedAt:              timestamppb.New(request.DeletedAt.Time),
			EquipmentRequestStatus: equipmentRequestStatus,
		},
	}, nil
}

// FormUpdatedEquipmentIDEvent is a function to create event about equipment id updating
func FormUpdatedEquipmentIDEvent(requestID uint64, equipmentID uint64) *EquipmentRequestEvent {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               UpdatedEquipmentID,
		CreatedAt:          time.Now(),
		EquipmentRequestID: requestID,
		Payload: &EquipmentRequestEventPayload{
			Id:          requestID,
			EquipmentId: equipmentID,
		},
	}
}

// FormUpdatedStatusEvent is a function to create event about request status updating
func FormUpdatedStatusEvent(requestID uint64, status EquipmentRequestStatus) (*EquipmentRequestEvent, error) {
	statusVal, ok := pb.EquipmentRequestStatus_value[string(status)]
	var equipmentRequestStatus pb.EquipmentRequestStatus

	if !ok {
		return nil, errors.New("unable to convert equipment request status")
	}

	equipmentRequestStatus = pb.EquipmentRequestStatus(statusVal)

	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               UpdatedStatus,
		CreatedAt:          time.Now(),
		EquipmentRequestID: requestID,
		Payload: &EquipmentRequestEventPayload{
			Id:                     requestID,
			EquipmentRequestStatus: equipmentRequestStatus,
		},
	}, nil
}

// FormRemovedEvent is a function to create event about equipment request removing
func FormRemovedEvent(equipmentRequestID uint64) *EquipmentRequestEvent {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               Removed,
		CreatedAt:          time.Now(),
		EquipmentRequestID: equipmentRequestID,
	}
}

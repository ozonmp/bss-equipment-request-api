package model

import (
	"database/sql"
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

func (es EventStatus) String() string {
	return string(es)
}

// EquipmentRequestEvent is a event of equipment request
type EquipmentRequestEvent struct {
	ID                 uint64            `db:"id"`
	Type               EventType         `db:"type"`
	Status             EventStatus       `db:"status"`
	CreatedAt          time.Time         `db:"created_at"`
	UpdatedAt          sql.NullTime      `db:"updated_at"`
	EquipmentRequestID uint64            `db:"equipment_request_id"`
	Payload            *EquipmentRequest `db:"payload"`
}

// FormCreatedEvent is a function to create event about equipment request creating
func FormCreatedEvent(request *EquipmentRequest) (*EquipmentRequestEvent, error) {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               Created,
		CreatedAt:          time.Now(),
		EquipmentRequestID: request.ID,
		Payload: &EquipmentRequest{
			ID:                     request.ID,
			EmployeeID:             request.EmployeeID,
			EquipmentID:            request.EquipmentID,
			CreatedAt:              request.CreatedAt,
			UpdatedAt:              request.UpdatedAt,
			DoneAt:                 request.DoneAt,
			DeletedAt:              request.DeletedAt,
			EquipmentRequestStatus: request.EquipmentRequestStatus,
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
		Payload: &EquipmentRequest{
			ID:          requestID,
			EquipmentID: equipmentID,
		},
	}
}

// FormUpdatedStatusEvent is a function to create event about request status updating
func FormUpdatedStatusEvent(requestID uint64, status EquipmentRequestStatus) (*EquipmentRequestEvent, error) {
	return &EquipmentRequestEvent{
		Status:             Unlocked,
		Type:               UpdatedStatus,
		CreatedAt:          time.Now(),
		EquipmentRequestID: requestID,
		Payload: &EquipmentRequest{
			ID:                     requestID,
			EquipmentRequestStatus: status,
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

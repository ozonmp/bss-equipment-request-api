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

func (es EventStatus) String() string {
	return string(es)
}

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
	ID                 uint64            `db:"id"`
	Type               EventType         `db:"type"`
	Status             EventStatus       `db:"status"`
	CreatedAt          time.Time         `db:"created_at"`
	UpdatedAt          sql.NullTime      `db:"updated_at"`
	EquipmentRequestID uint64            `db:"equipment_request_id"`
	Payload            *EquipmentRequest `db:"payload"`
}

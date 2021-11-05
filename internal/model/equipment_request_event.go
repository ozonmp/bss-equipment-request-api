package model

// EventType is a type of events
type EventType uint8

const (
	// Created is a "created item" type of events
	Created EventType = iota
	// Updated is a "updated item" type of events
	Updated
	// Removed is a "removed item" type of events
	Removed
)

// EventStatus is a status of events
type EventStatus uint8

const (
	// Deferred is a deferred status of events
	Deferred EventStatus = iota
	// Processed is a precessed status of events
	Processed
)

// EquipmentRequestEvent is a event of equipment request
type EquipmentRequestEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *EquipmentRequest
}

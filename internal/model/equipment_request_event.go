package model

type EventType uint8

const (
	Created EventType = iota
	Updated
	Removed
)

type EventStatus uint8

const (
	Deferred EventStatus = iota
	Processed
)

type EquipmentRequestEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *EquipmentRequest
}

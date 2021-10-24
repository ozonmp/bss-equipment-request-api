package model

type EquipmentRequest struct {
	Id            uint64
	EmployeeId    uint64
	EquipmentType string
	EquipmentId   uint64
	CreatedAt     string
	DoneAt        string
	Status        bool
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type EquipmentRequestEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *EquipmentRequest
}

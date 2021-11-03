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

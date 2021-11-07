package model

import "time"

// EquipmentRequest is a request for equipment
type EquipmentRequest struct {
	ID                       uint64                 `db:"id"`
	EmployeeID               uint64                 `db:"employee_id"`
	EquipmentID              uint64                 `db:"equipment_id"`
	CreatedAt                *time.Time             `db:"created_at"`
	DoneAt                   *time.Time             `db:"done_at"`
	EquipmentRequestStatusID EquipmentRequestStatus `db:"equipment_request_status_id"`
}

// EquipmentRequestStatus is a status of request for equipment
type EquipmentRequestStatus int

//EquipmentRequestStatuses
const (
	_ = iota
	// Do is a equipment request to do
	Do EquipmentRequestStatus = iota
	// InProgress is a equipment request in progress
	InProgress
	// Done is a done equipment request
	Done
	// Cancelled is a cancelled equipment request
	Cancelled
)

func (es EquipmentRequestStatus) String() string {
	return [...]string{"Do", "In Progress", "Done", "Cancelled"}[es]
}

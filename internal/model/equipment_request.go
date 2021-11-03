package model

import "time"

type EquipmentRequest struct {
	Id                       uint64                 `db:"id"`
	EmployeeId               uint64                 `db:"employee_id"`
	EquipmentId              uint64                 `db:"equipment_id"`
	CreatedAt                time.Time              `db:"created_at"`
	DoneAt                   time.Time              `db:"done_at"`
	EquipmentRequestStatusId EquipmentRequestStatus `db:"equipment_request_status_id"`
}
type EquipmentRequestStatus int

const (
	Do EquipmentRequestStatus = iota
	InProgress
	Done
	Cancelled
)

func (es EquipmentRequestStatus) String() string {
	return [...]string{"Do", "In Progress", "Done", "Cancelled"}[es]
}

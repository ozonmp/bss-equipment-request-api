package model

import (
	"database/sql"
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

// EquipmentRequest is a request for equipment
type EquipmentRequest struct {
	ID                     uint64                 `db:"id"`
	EmployeeID             uint64                 `db:"employee_id"`
	EquipmentID            uint64                 `db:"equipment_id"`
	CreatedAt              time.Time              `db:"created_at"`
	UpdatedAt              sql.NullTime           `db:"updated_at"`
	DeletedAt              sql.NullTime           `db:"deleted_at"`
	DoneAt                 sql.NullTime           `db:"done_at"`
	EquipmentRequestStatus EquipmentRequestStatus `db:"equipment_request_status"`
}

// Scan EquipmentRequestEventPayload
func (e *EquipmentRequest) Scan(src interface{}) (err error) {
	var eqp pb.EquipmentRequestPayload

	switch src.(type) {
	case string:
		err = protojson.Unmarshal([]byte(src.(string)), &eqp)
	case []byte:
		err = protojson.Unmarshal(src.([]byte), &eqp)
	default:
		return errors.New("incompatible type for EquipmentRequest")
	}

	if err != nil {
		return err
	}

	request := ConvertPbToEquipmentRequestPayload(&eqp)

	*e = *request

	return nil
}

// EquipmentRequestStatus is a status of request for equipment
type EquipmentRequestStatus string

//EquipmentRequestStatuses
const (
	// Do is an equipment request to do
	Do EquipmentRequestStatus = "EQUIPMENT_REQUEST_STATUS_DO"
	// InProgress is a equipment request in progress
	InProgress = "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS"
	// Done is a done equipment request
	Done = "EQUIPMENT_REQUEST_STATUS_DONE"
	// Cancelled is a cancelled equipment request
	Cancelled = "EQUIPMENT_REQUEST_STATUS_CANCELLED"
)

func (es EquipmentRequestStatus) String() string {
	return string(es)
}

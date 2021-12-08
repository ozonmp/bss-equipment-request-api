package model

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertEquipmentRequestPayloadToPb(t *testing.T) {
	tests := []struct {
		name  string
		input EquipmentRequest
		want  pb.EquipmentRequestPayload
	}{
		{

			name: "should_create_equipment_request",
			input: EquipmentRequest{
				ID:                     5,
				EmployeeID:             100,
				EquipmentID:            15,
				CreatedAt:              time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC),
				UpdatedAt:              sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				DeletedAt:              sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				DoneAt:                 sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_DONE"},
			want: pb.EquipmentRequestPayload{
				Id:                     5,
				EmployeeId:             100,
				EquipmentId:            15,
				CreatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				UpdatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				DeletedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				DoneAt:                 &timestamp.Timestamp{Seconds: 1640984399},
				EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_DONE"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertEquipmentRequestPayloadToPb(&tt.input)

			assert.Equal(t, tt.want.EmployeeId, got.EmployeeId)
			assert.Equal(t, tt.want.EquipmentId, got.EquipmentId)
			assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
			assert.Equal(t, tt.want.UpdatedAt, got.UpdatedAt)
			assert.Equal(t, tt.want.DeletedAt, got.DeletedAt)
			assert.Equal(t, tt.want.DoneAt, got.DoneAt)
			assert.Equal(t, tt.want.EquipmentRequestStatus, got.EquipmentRequestStatus)
		})
	}
}

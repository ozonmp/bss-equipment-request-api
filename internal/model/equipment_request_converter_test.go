package model

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertEquipmentRequestStatusToPb(t *testing.T) {
	tests := []struct {
		name  string
		input EquipmentRequest
		want  pb.EquipmentRequest
	}{
		{

			name: "should_create_max_employee_id",
			input: EquipmentRequest{
				ID:                     1,
				EmployeeID:             18446744073709551615,
				EquipmentID:            15,
				CreatedAt:              time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC),
				EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_DO"},
			want: pb.EquipmentRequest{
				Id:                     1,
				EmployeeId:             18446744073709551615,
				EquipmentId:            15,
				CreatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				UpdatedAt:              nil,
				DeletedAt:              nil,
				DoneAt:                 nil,
				EquipmentRequestStatus: 1},
		},
		{
			name: "should_create_with_status_default",
			input: EquipmentRequest{
				ID:                     2,
				EmployeeID:             1,
				EquipmentID:            2,
				CreatedAt:              time.Date(2022, 01, 01, 23, 59, 59, 0, time.UTC),
				UpdatedAt:              sql.NullTime{time.Date(2022, 02, 01, 23, 59, 59, 0, time.UTC), true},
				EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_UNSPECIFIED"},
			want: pb.EquipmentRequest{
				Id:                     2,
				EmployeeId:             1,
				EquipmentId:            2,
				CreatedAt:              &timestamp.Timestamp{Seconds: 1641081599},
				UpdatedAt:              &timestamp.Timestamp{Seconds: 1643759999},
				DeletedAt:              nil,
				DoneAt:                 nil,
				EquipmentRequestStatus: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertEquipmentRequestToPb(&tt.input)

			assert.NoError(t, err)
			assert.Equal(t, tt.want.Id, got.Id)
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

func TestConvertRepeatedEquipmentRequestsToPb2(t *testing.T) {
	tests := []struct {
		name  string
		input []EquipmentRequest
		want  []pb.EquipmentRequest
	}{
		{

			name: "should_multi_create",
			input: []EquipmentRequest{
				{
					ID:                     1,
					EmployeeID:             18446744073709551615,
					EquipmentID:            15,
					CreatedAt:              time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC),
					EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_UNSPECIFIED"},
				{
					ID:                     2,
					EmployeeID:             1,
					EquipmentID:            2,
					CreatedAt:              time.Date(2022, 01, 01, 23, 59, 59, 0, time.UTC),
					UpdatedAt:              sql.NullTime{time.Date(2022, 02, 01, 23, 59, 59, 0, time.UTC), true},
					EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_DO"}},
			want: []pb.EquipmentRequest{
				{
					Id:                     1,
					EmployeeId:             18446744073709551615,
					EquipmentId:            15,
					CreatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
					UpdatedAt:              nil,
					DeletedAt:              nil,
					DoneAt:                 nil,
					EquipmentRequestStatus: 0},
				{
					Id:                     2,
					EmployeeId:             1,
					EquipmentId:            2,
					CreatedAt:              &timestamp.Timestamp{Seconds: 1641081599},
					UpdatedAt:              &timestamp.Timestamp{Seconds: 1643759999},
					DeletedAt:              nil,
					DoneAt:                 nil,
					EquipmentRequestStatus: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertRepeatedEquipmentRequestsToPb(tt.input)

			assert.NoError(t, err)

			for equipment := range got {
				assert.Equal(t, tt.want[equipment].Id, got[equipment].Id)
				assert.Equal(t, tt.want[equipment].EmployeeId, got[equipment].EmployeeId)
				assert.Equal(t, tt.want[equipment].EquipmentId, got[equipment].EquipmentId)
				assert.Equal(t, tt.want[equipment].CreatedAt, got[equipment].CreatedAt)
				assert.Equal(t, tt.want[equipment].UpdatedAt, got[equipment].UpdatedAt)
				assert.Equal(t, tt.want[equipment].DeletedAt, got[equipment].DeletedAt)
				assert.Equal(t, tt.want[equipment].DoneAt, got[equipment].DoneAt)
				assert.Equal(t, tt.want[equipment].EquipmentRequestStatus, got[equipment].EquipmentRequestStatus)
			}

		})
	}
}

func TestConvertPbToEquipmentRequest(t *testing.T) {
	tests := []struct {
		name  string
		input pb.EquipmentRequest
		want  EquipmentRequest
	}{
		{

			name: "should_create_equipment_request",
			input: pb.EquipmentRequest{
				Id:                     5,
				EmployeeId:             100,
				EquipmentId:            15,
				CreatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				UpdatedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				DeletedAt:              &timestamp.Timestamp{Seconds: 1640984399},
				DoneAt:                 &timestamp.Timestamp{Seconds: 1640984399},
				EquipmentRequestStatus: 3},
			want: EquipmentRequest{
				ID:                     5,
				EmployeeID:             100,
				EquipmentID:            15,
				CreatedAt:              time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC),
				UpdatedAt:              sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				DeletedAt:              sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				DoneAt:                 sql.NullTime{time.Date(2021, 12, 31, 20, 59, 59, 0, time.UTC), true},
				EquipmentRequestStatus: "EQUIPMENT_REQUEST_STATUS_DONE"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertPbToEquipmentRequest(&tt.input)

			assert.NoError(t, err)
			assert.Equal(t, tt.want.EmployeeID, got.EmployeeID)
			assert.Equal(t, tt.want.EquipmentID, got.EquipmentID)
			assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
			assert.Equal(t, tt.want.UpdatedAt, got.UpdatedAt)
			assert.Equal(t, tt.want.DeletedAt, got.DeletedAt)
			assert.Equal(t, tt.want.DoneAt, got.DoneAt)
			assert.Equal(t, tt.want.EquipmentRequestStatus, got.EquipmentRequestStatus)
		})
	}
}

func TestConvertEquipmentRequestStatusToPb2(t *testing.T) {
	tests := []struct {
		name  string
		input EquipmentRequestStatus
		want  pb.EquipmentRequestStatus
	}{
		{
			name:  "should_get_status_0",
			input: "EQUIPMENT_REQUEST_STATUS_UNSPECIFIED",
			want:  0,
		},
		{
			name:  "should_get_status_1",
			input: "EQUIPMENT_REQUEST_STATUS_DO",
			want:  1,
		},
		{
			name:  "should_get_status_2",
			input: "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS",
			want:  2,
		},
		{
			name:  "should_get_status_3",
			input: "EQUIPMENT_REQUEST_STATUS_DONE",
			want:  3,
		},
		{
			name:  "should_get_status_4",
			input: "EQUIPMENT_REQUEST_STATUS_CANCELLED",
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ConvertEquipmentRequestStatusToPb(tt.input)

			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}

func TestConvertPbEquipmentRequestStatus2(t *testing.T) {
	tests := []struct {
		name  string
		input pb.EquipmentRequestStatus
		want  EquipmentRequestStatus
	}{
		{
			name:  "should_convert_status_UNSPECIFIED",
			input: 0,
			want:  "EQUIPMENT_REQUEST_STATUS_UNSPECIFIED",
		},
		{
			name:  "should_convert_status_DO",
			input: 1,
			want:  "EQUIPMENT_REQUEST_STATUS_DO",
		},
		{
			name:  "should_convert_status_IN_PROGRESS",
			input: 2,
			want:  "EQUIPMENT_REQUEST_STATUS_IN_PROGRESS",
		},
		{
			name:  "should_convert_status_DONE",
			input: 3,
			want:  "EQUIPMENT_REQUEST_STATUS_DONE",
		},
		{
			name:  "should_convert_status_CANCELLED",
			input: 4,
			want:  "EQUIPMENT_REQUEST_STATUS_CANCELLED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ConvertPbEquipmentRequestStatus(tt.input)

			assert.NoError(t, err)
			assert.Equal(t, &tt.want, got)
		})
	}
}

package model

import (
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
)

var (
	// ErrUnableToConvertEquipmentRequestStatus is a unable to convert model error
	ErrUnableToConvertEquipmentRequestStatus = errors.New("unable to convert equipment request status")
)

//ConvertCreatePbMessageToEquipmentRequest - convert protobuf CreateEquipmentRequestV1Request message to EquipmentRequest
func ConvertCreatePbMessageToEquipmentRequest(equipmentRequest *pb.CreateEquipmentRequestV1Request) (*EquipmentRequest, error) {
	equipmentRequestStatus, err := ConvertPbEquipmentRequestStatus(equipmentRequest.EquipmentRequestStatus)

	if err != nil {
		return nil, err
	}

	updatedAtTime := ConvertPbTimeToNullableTime(equipmentRequest.UpdatedAt)
	doneAtTime := ConvertPbTimeToNullableTime(equipmentRequest.DoneAt)
	deletedAtTime := ConvertPbTimeToNullableTime(equipmentRequest.DeletedAt)

	return &EquipmentRequest{
		EmployeeID:             equipmentRequest.EmployeeId,
		EquipmentID:            equipmentRequest.EquipmentId,
		CreatedAt:              equipmentRequest.CreatedAt.AsTime(),
		UpdatedAt:              updatedAtTime,
		DoneAt:                 doneAtTime,
		DeletedAt:              deletedAtTime,
		EquipmentRequestStatus: *equipmentRequestStatus,
	}, nil
}

//ConvertPbEquipmentRequestStatus - convert protobuf EquipmentRequestStatus enum to EquipmentRequestStatus
func ConvertPbEquipmentRequestStatus(equipmentRequestStatus pb.EquipmentRequestStatus) (*EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_name[int32(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestModelStatus := EquipmentRequestStatus(status)

	return &equipmentRequestModelStatus, nil
}

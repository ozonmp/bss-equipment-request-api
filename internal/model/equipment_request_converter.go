package model

import (
	"errors"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// ErrUnableToConvertEquipmentRequestStatus is a unable to convert model error
	ErrUnableToConvertEquipmentRequestStatus = errors.New("unable to convert equipment request status")
)

//ConvertEquipmentRequestToPb - convert EquipmentRequest to protobuf EquipmentRequest message
func ConvertEquipmentRequestToPb(equipmentRequest *EquipmentRequest) (*pb.EquipmentRequest, error) {
	equipmentRequestStatus, err := ConvertEquipmentRequestStatusToPb(equipmentRequest.EquipmentRequestStatus)

	if err != nil {
		return nil, err
	}

	return &pb.EquipmentRequest{
		Id:                     equipmentRequest.ID,
		EmployeeId:             equipmentRequest.EmployeeID,
		EquipmentId:            equipmentRequest.EquipmentID,
		CreatedAt:              timestamppb.New(equipmentRequest.CreatedAt),
		UpdatedAt:              timestamppb.New(equipmentRequest.UpdatedAt.Time),
		DoneAt:                 timestamppb.New(equipmentRequest.DoneAt.Time),
		DeletedAt:              timestamppb.New(equipmentRequest.DeletedAt.Time),
		EquipmentRequestStatus: *equipmentRequestStatus,
	}, nil
}

//ConvertRepeatedEquipmentRequestsToPb - convert slice of EquipmentRequest to slice of protobuf EquipmentRequest messages
func ConvertRepeatedEquipmentRequestsToPb(equipmentRequests []EquipmentRequest) ([]*pb.EquipmentRequest, error) {
	var equipmentRequestsPb []*pb.EquipmentRequest

	for i := range equipmentRequests {
		equipmentRequest, err := ConvertEquipmentRequestToPb(&equipmentRequests[i])
		if err != nil {
			return nil, err
		}
		equipmentRequestsPb = append(equipmentRequestsPb, equipmentRequest)
	}

	return equipmentRequestsPb, nil
}

//ConvertPbToEquipmentRequest - convert protobuf EquipmentRequest message to EquipmentRequest
func ConvertPbToEquipmentRequest(equipmentRequest *pb.EquipmentRequest) (*EquipmentRequest, error) {
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

//ConvertEquipmentRequestStatusToPb - convert EquipmentRequestStatus to protobuf EquipmentRequestStatus enum
func ConvertEquipmentRequestStatusToPb(equipmentRequestStatus EquipmentRequestStatus) (*pb.EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_value[string(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestPbStatus := pb.EquipmentRequestStatus(status)

	return &equipmentRequestPbStatus, nil
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

package model

import (
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//ConvertEquipmentRequestPayloadToPb - convert EquipmentRequest to protobuf EquipmentRequestPayload
func ConvertEquipmentRequestPayloadToPb(equipmentRequest *EquipmentRequest) *pb.EquipmentRequestPayload {
	payload := &pb.EquipmentRequestPayload{
		Id:                     equipmentRequest.ID,
		EmployeeId:             equipmentRequest.EmployeeID,
		EquipmentId:            equipmentRequest.EquipmentID,
		EquipmentRequestStatus: equipmentRequest.EquipmentRequestStatus.String(),
	}

	if !equipmentRequest.CreatedAt.IsZero() {
		payload.CreatedAt = timestamppb.New(equipmentRequest.CreatedAt)
	}

	if !equipmentRequest.UpdatedAt.Time.IsZero() {
		payload.UpdatedAt = timestamppb.New(equipmentRequest.UpdatedAt.Time)
	}

	if !equipmentRequest.DoneAt.Time.IsZero() {
		payload.DoneAt = timestamppb.New(equipmentRequest.DoneAt.Time)
	}

	if !equipmentRequest.DeletedAt.Time.IsZero() {
		payload.DeletedAt = timestamppb.New(equipmentRequest.DeletedAt.Time)
	}

	return payload
}

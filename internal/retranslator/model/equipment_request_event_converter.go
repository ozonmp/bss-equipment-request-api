package model

import (
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ConvertPbToEquipmentRequestPayload - convert protobuf EquipmentRequestPayload message to EquipmentRequest
func ConvertPbToEquipmentRequestPayload(equipmentRequestPayload *pb.EquipmentRequestPayload) *EquipmentRequest {
	return &EquipmentRequest{
		ID:                     equipmentRequestPayload.Id,
		EmployeeID:             equipmentRequestPayload.EmployeeId,
		EquipmentID:            equipmentRequestPayload.EquipmentId,
		CreatedAt:              equipmentRequestPayload.CreatedAt.AsTime(),
		UpdatedAt:              ConvertPbTimeToNullableTime(equipmentRequestPayload.UpdatedAt),
		DoneAt:                 ConvertPbTimeToNullableTime(equipmentRequestPayload.DoneAt),
		DeletedAt:              ConvertPbTimeToNullableTime(equipmentRequestPayload.DeletedAt),
		EquipmentRequestStatus: EquipmentRequestStatus(equipmentRequestPayload.EquipmentRequestStatus),
	}
}

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

//ConvertPbToEquipmentRequestEvent - convert protobuf EquipmentRequestEvent message to EquipmentRequestEvent
func ConvertPbToEquipmentRequestEvent(equipmentRequestEvent *pb.EquipmentRequestEvent) *EquipmentRequestEvent {
	event := &EquipmentRequestEvent{
		ID:                 equipmentRequestEvent.Id,
		Type:               EventType(equipmentRequestEvent.Type),
		Status:             EventStatus(equipmentRequestEvent.Status),
		EquipmentRequestID: equipmentRequestEvent.EquipmentRequestId,
		CreatedAt:          equipmentRequestEvent.CreatedAt.AsTime(),
		UpdatedAt:          ConvertPbTimeToNullableTime(equipmentRequestEvent.UpdatedAt),
	}

	if equipmentRequestEvent.Payload != nil {
		payload := ConvertPbToEquipmentRequestPayload(equipmentRequestEvent.Payload)

		event.Payload = payload
	}

	return event
}

// ConvertEquipmentRequestEventToPb is a function to convert EquipmentRequestEvent to pb.EquipmentRequestEvent
func ConvertEquipmentRequestEventToPb(equipmentRequestEvent *EquipmentRequestEvent) (*pb.EquipmentRequestEvent, error) {
	event := &pb.EquipmentRequestEvent{
		Id:                 equipmentRequestEvent.ID,
		EquipmentRequestId: equipmentRequestEvent.EquipmentRequestID,
		Type:               equipmentRequestEvent.Type.String(),
		Status:             equipmentRequestEvent.Status.String(),
		CreatedAt:          timestamppb.New(equipmentRequestEvent.CreatedAt),
		UpdatedAt:          timestamppb.New(equipmentRequestEvent.UpdatedAt.Time),
	}

	if equipmentRequestEvent.Payload != nil {
		payload := ConvertEquipmentRequestPayloadToPb(equipmentRequestEvent.Payload)

		event.Payload = payload
	}

	return event, nil
}

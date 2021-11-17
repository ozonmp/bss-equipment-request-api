package api

import (
	"database/sql"
	"errors"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	createEquipmentRequestV1LogTag            = "CreateEquipmentRequestV1"
	describeEquipmentRequestV1LogTag          = "DescribeEquipmentRequestV1"
	listEquipmentRequestV1LogTag              = "ListEquipmentRequestV1"
	removeEquipmentRequestV1LogTag            = "RemoveEquipmentRequestV1"
	updateEquipmentIDEquipmentRequestV1LogTag = "UpdateEquipmentIDEquipmentRequestV1"
	updateStatusEquipmentRequestV1LogTag      = "UpdateStatusEquipmentRequestV1"
)

var (
	// ErrUnableToConvertEquipmentRequestStatus is a unable to convert model error
	ErrUnableToConvertEquipmentRequestStatus = errors.New("unable to convert equipment request status")
)

type equipmentRequestAPI struct {
	pb.UnimplementedBssEquipmentRequestApiServiceServer
	equipmentRequestService equipment_request.ServiceInterface
}

// NewEquipmentRequestAPI returns api of bss-equipment-request-api service
func NewEquipmentRequestAPI(equipmentRequestService equipment_request.ServiceInterface) pb.BssEquipmentRequestApiServiceServer {
	return &equipmentRequestAPI{equipmentRequestService: equipmentRequestService}
}

func (o *equipmentRequestAPI) convertEquipmentRequestToPb(equipmentRequest *model.EquipmentRequest) (*pb.EquipmentRequest, error) {
	equipmentRequestStatus, err := o.convertEquipmentRequestStatusToPb(equipmentRequest.EquipmentRequestStatus)

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

func (o *equipmentRequestAPI) convertPbToEquipmentRequest(equipmentRequest *pb.EquipmentRequest) (*model.EquipmentRequest, error) {
	equipmentRequestStatus, err := o.convertPbEquipmentRequestStatus(equipmentRequest.EquipmentRequestStatus)

	if err != nil {
		return nil, err
	}

	updatedAtTime := o.convertPbTimeToNullableTime(equipmentRequest.UpdatedAt)
	doneAtTime := o.convertPbTimeToNullableTime(equipmentRequest.DoneAt)
	deletedAtTime := o.convertPbTimeToNullableTime(equipmentRequest.DeletedAt)

	return &model.EquipmentRequest{
		EmployeeID:             equipmentRequest.EmployeeId,
		EquipmentID:            equipmentRequest.EquipmentId,
		CreatedAt:              equipmentRequest.CreatedAt.AsTime(),
		UpdatedAt:              updatedAtTime,
		DoneAt:                 doneAtTime,
		DeletedAt:              deletedAtTime,
		EquipmentRequestStatus: *equipmentRequestStatus,
	}, nil
}

func (o *equipmentRequestAPI) convertPbTimeToNullableTime(pbTime *timestamppb.Timestamp) sql.NullTime {
	var deletedAtTime sql.NullTime
	if pbTime != nil {
		deletedAtTime = sql.NullTime{Time: pbTime.AsTime(), Valid: true}
	}

	return deletedAtTime
}

func (o *equipmentRequestAPI) convertRepeatedEquipmentRequestsToPb(equipmentRequests []model.EquipmentRequest) ([]*pb.EquipmentRequest, error) {
	var equipmentRequestsPb []*pb.EquipmentRequest

	for i := range equipmentRequests {
		equipmentRequest, err := o.convertEquipmentRequestToPb(&equipmentRequests[i])
		if err != nil {
			return nil, err
		}
		equipmentRequestsPb = append(equipmentRequestsPb, equipmentRequest)
	}

	return equipmentRequestsPb, nil
}

func (o *equipmentRequestAPI) convertEquipmentRequestStatusToPb(equipmentRequestStatus model.EquipmentRequestStatus) (*pb.EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_value[string(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestPbStatus := pb.EquipmentRequestStatus(status)

	return &equipmentRequestPbStatus, nil
}

func (o *equipmentRequestAPI) convertPbEquipmentRequestStatus(equipmentRequestStatus pb.EquipmentRequestStatus) (*model.EquipmentRequestStatus, error) {
	status, ok := pb.EquipmentRequestStatus_name[int32(equipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestModelStatus := model.EquipmentRequestStatus(status)

	return &equipmentRequestModelStatus, nil
}

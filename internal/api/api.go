package api

import (
	"database/sql"
	"errors"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	totalEquipmentRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_equipment_request_api_not_found_total",
		Help: "Total number of equipment requests that were not found",
	})

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
	status, ok := pb.EquipmentRequestStatus_value[string(equipmentRequest.EquipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestStatus := pb.EquipmentRequestStatus(status)

	return &pb.EquipmentRequest{
		Id:                     equipmentRequest.ID,
		EmployeeId:             equipmentRequest.EmployeeID,
		EquipmentId:            equipmentRequest.EquipmentID,
		CreatedAt:              timestamppb.New(equipmentRequest.CreatedAt),
		UpdatedAt:              timestamppb.New(equipmentRequest.UpdatedAt.Time),
		DoneAt:                 timestamppb.New(equipmentRequest.DoneAt.Time),
		DeletedAt:              timestamppb.New(equipmentRequest.DeletedAt.Time),
		EquipmentRequestStatus: equipmentRequestStatus,
	}, nil
}

func (o *equipmentRequestAPI) convertPbToEquipmentRequest(equipmentRequest *pb.EquipmentRequest) (*model.EquipmentRequest, error) {
	status, ok := pb.EquipmentRequestStatus_name[int32(equipmentRequest.EquipmentRequestStatus)]

	if !ok {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	equipmentRequestStatus := model.EquipmentRequestStatus(status)

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
		EquipmentRequestStatus: equipmentRequestStatus,
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

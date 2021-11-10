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
	"time"
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
	var equipmentRequestStatus pb.EquipmentRequestStatus

	if ok {
		equipmentRequestStatus = pb.EquipmentRequestStatus(status)
	} else {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	return &pb.EquipmentRequest{
		Id:                     equipmentRequest.ID,
		EmployeeId:             equipmentRequest.EmployeeID,
		EquipmentId:            equipmentRequest.EquipmentID,
		CreatedAt:              timestamppb.New(equipmentRequest.CreatedAt),
		UpdatedAt:              timestamppb.New(equipmentRequest.UpdatedAt),
		DoneAt:                 timestamppb.New(equipmentRequest.DoneAt.Time),
		DeletedAt:              timestamppb.New(equipmentRequest.DeletedAt.Time),
		EquipmentRequestStatus: equipmentRequestStatus,
	}, nil
}

func (o *equipmentRequestAPI) convertPbToEquipmentRequest(equipmentRequest *pb.EquipmentRequest) (*model.EquipmentRequest, error) {
	status, ok := pb.EquipmentRequestStatus_name[int32(equipmentRequest.EquipmentRequestStatus)]
	var equipmentRequestStatus model.EquipmentRequestStatus

	if ok {
		equipmentRequestStatus = model.EquipmentRequestStatus(status)
	} else {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	var createdAtTime time.Time
	if equipmentRequest.CreatedAt == nil {
		createdAtTime = time.Now()
	} else {
		createdAtTime = equipmentRequest.CreatedAt.AsTime()
	}

	var updatedAtTime time.Time
	if equipmentRequest.UpdatedAt == nil {
		updatedAtTime = time.Now()
	} else {
		updatedAtTime = equipmentRequest.UpdatedAt.AsTime()
	}

	var doneAtTime sql.NullTime
	if equipmentRequest.DoneAt == nil {
		doneAtTime = sql.NullTime{Time: time.Time{}, Valid: false}
	} else {
		doneAtTime = sql.NullTime{Time: equipmentRequest.DoneAt.AsTime(), Valid: true}
	}

	var deletedAtTime sql.NullTime
	if equipmentRequest.DeletedAt == nil {
		deletedAtTime = sql.NullTime{Time: time.Time{}, Valid: false}
	} else {
		deletedAtTime = sql.NullTime{Time: equipmentRequest.DeletedAt.AsTime(), Valid: true}
	}

	return &model.EquipmentRequest{
		EmployeeID:             equipmentRequest.EmployeeId,
		EquipmentID:            equipmentRequest.EquipmentId,
		CreatedAt:              createdAtTime,
		UpdatedAt:              updatedAtTime,
		DoneAt:                 doneAtTime,
		DeletedAt:              deletedAtTime,
		EquipmentRequestStatus: equipmentRequestStatus,
	}, nil
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

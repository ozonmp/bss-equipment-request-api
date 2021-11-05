package api

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalEquipmentRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_equipment_request_api_not_found_total",
		Help: "Total number of equipment requests that were not found",
	})
)

type equipmentRequestAPI struct {
	pb.UnimplementedBssEquipmentRequestApiServiceServer
	equipmentRequestService equipment_request.ServiceInterface
}

// NewEquipmentRequestAPI returns api of bss-equipment-request-api service
func NewEquipmentRequestAPI(equipmentRequestService equipment_request.ServiceInterface) pb.BssEquipmentRequestApiServiceServer {
	return &equipmentRequestAPI{equipmentRequestService: equipmentRequestService}
}

func (o *equipmentRequestAPI) convertEquipmentRequestToPb(equipmentRequest *model.EquipmentRequest) *pb.EquipmentRequest {
	return &pb.EquipmentRequest{
		Id:                       equipmentRequest.ID,
		EmployeeId:               equipmentRequest.EmployeeID,
		EquipmentId:              equipmentRequest.EquipmentID,
		CreatedAt:                timestamppb.New(*equipmentRequest.CreatedAt),
		DoneAt:                   timestamppb.New(*equipmentRequest.DoneAt),
		EquipmentRequestStatusId: uint64(equipmentRequest.EquipmentRequestStatusID),
	}
}

func (o *equipmentRequestAPI) convertRepeatedEquipmentRequestsToPb(equipmentRequests []model.EquipmentRequest) []*pb.EquipmentRequest {
	var equipmentRequestsPb []*pb.EquipmentRequest

	for i := range equipmentRequests {
		equipmentRequestsPb = append(equipmentRequestsPb, o.convertEquipmentRequestToPb(&equipmentRequests[i]))
	}

	return equipmentRequestsPb
}

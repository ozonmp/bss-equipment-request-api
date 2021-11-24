package api

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
)

const (
	createEquipmentRequestV1LogTag            = "CreateEquipmentRequestV1"
	describeEquipmentRequestV1LogTag          = "DescribeEquipmentRequestV1"
	listEquipmentRequestV1LogTag              = "ListEquipmentRequestV1"
	removeEquipmentRequestV1LogTag            = "RemoveEquipmentRequestV1"
	updateEquipmentIDEquipmentRequestV1LogTag = "UpdateEquipmentIDEquipmentRequestV1"
	updateStatusEquipmentRequestV1LogTag      = "UpdateStatusEquipmentRequestV1"
)

type equipmentRequestAPI struct {
	pb.UnimplementedBssEquipmentRequestApiServiceServer
	equipmentRequestService equipment_request.ServiceInterface
}

// NewEquipmentRequestAPI returns api of bss-equipment-request-api service
func NewEquipmentRequestAPI(equipmentRequestService equipment_request.ServiceInterface) pb.BssEquipmentRequestApiServiceServer {
	return &equipmentRequestAPI{equipmentRequestService: equipmentRequestService}
}

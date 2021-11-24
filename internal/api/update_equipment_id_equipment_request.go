package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) UpdateEquipmentIDEquipmentRequestV1(
	ctx context.Context,
	req *pb.UpdateEquipmentIDEquipmentRequestV1Request,
) (*pb.UpdateEquipmentIDEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid argument", updateEquipmentIDEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.equipmentRequestService.UpdateEquipmentIDEquipmentRequest(ctx, req.EquipmentRequestId, req.EquipmentId)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.UpdateEquipmentIDEquipmentRequest failed", updateEquipmentIDEquipmentRequestV1LogTag),
			"err", err,
			"equipmentRequestId", req.EquipmentId,
			"equipmentId", req.EquipmentId,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.UpdateEquipmentIDEquipmentRequest failed", updateEquipmentIDEquipmentRequestV1LogTag),
			"err", "unable to update equipment id of equipment request, no rows affected",
			"equipmentRequestId", req.EquipmentId,
			"equipmentId", req.EquipmentId,
		)

		return nil, status.Error(codes.Internal, "unable to update equipment id of equipment request")
	}

	logger.Info(ctx, fmt.Sprintf("%s: success", updateEquipmentIDEquipmentRequestV1LogTag))

	return &pb.UpdateEquipmentIDEquipmentRequestV1Response{
		Updated: result,
	}, nil
}

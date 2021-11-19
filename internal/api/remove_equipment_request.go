package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) RemoveEquipmentRequestV1(
	ctx context.Context,
	req *pb.RemoveEquipmentRequestV1Request,
) (*pb.RemoveEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid argument", removeEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.equipmentRequestService.RemoveEquipmentRequest(ctx, req.EquipmentRequestId)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.RemoveEquipmentRequest failed", removeEquipmentRequestV1LogTag),
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.RemoveEquipmentRequest failed", removeEquipmentRequestV1LogTag),
			"err", "unable to remove equipment request, no rows affected",
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, "unable to remove equipment request")
	}

	logger.Info(ctx, fmt.Sprintf("%s: success", removeEquipmentRequestV1LogTag))

	return &pb.RemoveEquipmentRequestV1Response{
		Removed: result,
	}, nil
}

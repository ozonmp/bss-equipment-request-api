package api

import (
	"context"
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
		logger.ErrorKV(ctx, removeEquipmentRequestV1LogTag+": invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := o.equipmentRequestService.RemoveEquipmentRequest(ctx, req.EquipmentRequestId)

	if err != nil {
		logger.ErrorKV(ctx, removeEquipmentRequestV1LogTag+": equipmentRequestService.RemoveEquipmentRequest failed",
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.DebugKV(ctx, removeEquipmentRequestV1LogTag+": equipmentRequestService.RemoveEquipmentRequest failed",
			"err", "unable to remove equipment request, no rows affected",
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, "unable to remove equipment request")
	}

	logger.InfoKV(ctx, removeEquipmentRequestV1LogTag, "success")

	return &pb.RemoveEquipmentRequestV1Response{
		Removed: result,
	}, nil
}

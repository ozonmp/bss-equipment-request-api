package api

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) ListEquipmentRequestV1(
	ctx context.Context,
	req *pb.ListEquipmentRequestV1Request,
) (*pb.ListEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, listEquipmentRequestV1LogTag+": invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	equipmentRequests, err := o.equipmentRequestService.ListEquipmentRequest(ctx, req.Limit, req.Offset)

	if err != nil {
		logger.ErrorKV(ctx, listEquipmentRequestV1LogTag+": equipmentRequestService.ListEquipmentRequest failed",
			"err", err,
			"limit", req.Limit,
			"offset", req.Offset,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequests == nil {
		logger.DebugKV(ctx, listEquipmentRequestV1LogTag+": equipmentRequestService.ListEquipmentRequest failed",
			"err", "unable to get list of equipment requests",
			"limit", req.Limit,
			"offset", req.Offset,
		)

		return nil, status.Error(codes.NotFound, "unable to get list of equipment requests")
	}

	equipmentRequestPb, err := o.convertRepeatedEquipmentRequestsToPb(equipmentRequests)

	if err != nil {
		logger.ErrorKV(ctx, listEquipmentRequestV1LogTag+": unable to convert list of EquipmentRequests to Pb message",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.InfoKV(ctx, listEquipmentRequestV1LogTag, "success")

	return &pb.ListEquipmentRequestV1Response{
		Items: equipmentRequestPb,
	}, nil
}

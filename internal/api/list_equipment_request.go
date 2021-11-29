package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) ListEquipmentRequestV1(
	ctx context.Context,
	req *pb.ListEquipmentRequestV1Request,
) (*pb.ListEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid argument", listEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	equipmentRequests, err := o.equipmentRequestService.ListEquipmentRequest(ctx, req.Limit, req.Offset)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.ListEquipmentRequest failed", listEquipmentRequestV1LogTag),
			"err", err,
			"limit", req.Limit,
			"offset", req.Offset,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequests == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.ListEquipmentRequest failed", listEquipmentRequestV1LogTag),
			"err", "unable to get list of equipment requests",
			"limit", req.Limit,
			"offset", req.Offset,
		)

		return nil, status.Error(codes.NotFound, "unable to get list of equipment requests")
	}

	equipmentRequestPb, err := model.ConvertRepeatedEquipmentRequestsToPb(equipmentRequests)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: unable to convert list of EquipmentRequests to Pb message", listEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.Info(ctx, fmt.Sprintf("%s: success", listEquipmentRequestV1LogTag))

	return &pb.ListEquipmentRequestV1Response{
		Items: equipmentRequestPb,
	}, nil
}

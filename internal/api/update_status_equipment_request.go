package api

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) UpdateStatusEquipmentRequestV1(
	ctx context.Context,
	req *pb.UpdateStatusEquipmentRequestV1Request,
) (*pb.UpdateStatusEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, updateStatusEquipmentRequestV1LogTag+": invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequestStatus, err := o.convertPbEquipmentRequestStatus(req.EquipmentRequestStatus)

	if err != nil {
		logger.ErrorKV(ctx, updateStatusEquipmentRequestV1LogTag+": unable to convert Pb EquipmentRequestStatus to EquipmentRequestStatus",
			"err", err,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := o.equipmentRequestService.UpdateStatusEquipmentRequest(ctx, req.EquipmentRequestId, *equipmentRequestStatus)

	if err != nil {
		logger.ErrorKV(ctx, updateStatusEquipmentRequestV1LogTag+": equipmentRequestService.UpdateStatusEquipmentRequest failed",
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.DebugKV(ctx, updateStatusEquipmentRequestV1LogTag+": failed",
			"err", "unable to update update status of equipment request, no rows affected",
			"equipmentRequestId", req.EquipmentRequestId,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, "unable to update update status of equipment request")
	}

	logger.InfoKV(ctx, updateStatusEquipmentRequestV1LogTag, "success")

	return &pb.UpdateStatusEquipmentRequestV1Response{
		Updated: result,
	}, nil
}

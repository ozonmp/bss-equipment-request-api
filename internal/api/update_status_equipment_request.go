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

func (o *equipmentRequestAPI) UpdateStatusEquipmentRequestV1(
	ctx context.Context,
	req *pb.UpdateStatusEquipmentRequestV1Request,
) (*pb.UpdateStatusEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid argument", updateStatusEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequestStatus, err := model.ConvertPbEquipmentRequestStatus(req.EquipmentRequestStatus)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: unable to convert Pb EquipmentRequestStatus to EquipmentRequestStatus", updateStatusEquipmentRequestV1LogTag),
			"err", err,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := o.equipmentRequestService.UpdateStatusEquipmentRequest(ctx, req.EquipmentRequestId, *equipmentRequestStatus)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.UpdateStatusEquipmentRequest failed", updateStatusEquipmentRequestV1LogTag),
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.UpdateStatusEquipmentRequest failed", updateStatusEquipmentRequestV1LogTag),
			"err", "unable to update update status of equipment request, no rows affected",
			"equipmentRequestId", req.EquipmentRequestId,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, "unable to update update status of equipment request")
	}

	logger.Info(ctx, fmt.Sprintf("%s: success", updateStatusEquipmentRequestV1LogTag))

	return &pb.UpdateStatusEquipmentRequestV1Response{
		Updated: result,
	}, nil
}

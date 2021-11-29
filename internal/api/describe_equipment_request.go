package api

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/metrics"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) DescribeEquipmentRequestV1(
	ctx context.Context,
	req *pb.DescribeEquipmentRequestV1Request,
) (*pb.DescribeEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: invalid argument", describeEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequest, err := o.equipmentRequestService.DescribeEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.DescribeEquipmentRequest failed", describeEquipmentRequestV1LogTag),
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequest == nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.DescribeEquipmentRequest failed", describeEquipmentRequestV1LogTag),
			"err", "equipment request not found",
			"equipmentRequestId", req.EquipmentRequestId,
		)

		metrics.IncTotalEquipmentRequestNotFound()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	equipmentRequestPb, err := model.ConvertEquipmentRequestToPb(equipmentRequest)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: unable to convert EquipmentRequest to Pb message", describeEquipmentRequestV1LogTag),
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.Info(ctx, fmt.Sprintf("%s: success", describeEquipmentRequestV1LogTag))

	return &pb.DescribeEquipmentRequestV1Response{
		EquipmentRequest: equipmentRequestPb,
	}, nil
}

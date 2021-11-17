package api

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/metrics"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) DescribeEquipmentRequestV1(
	ctx context.Context,
	req *pb.DescribeEquipmentRequestV1Request,
) (*pb.DescribeEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, describeEquipmentRequestV1LogTag+": invalid argument",
			"err", err,
		)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequest, err := o.equipmentRequestService.DescribeEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		logger.ErrorKV(ctx, describeEquipmentRequestV1LogTag+": equipmentRequestService.DescribeEquipmentRequest failed",
			"err", err,
			"equipmentRequestId", req.EquipmentRequestId,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequest == nil {
		logger.DebugKV(ctx, describeEquipmentRequestV1LogTag+": equipmentRequestService.DescribeEquipmentRequest failed",
			"err", "equipment request not found",
			"equipmentRequestId", req.EquipmentRequestId,
		)

		metrics.IncTotalEquipmentRequestNotFound()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	equipmentRequestPb, err := o.convertEquipmentRequestToPb(equipmentRequest)

	if err != nil {
		logger.ErrorKV(ctx, describeEquipmentRequestV1LogTag+": unable to convert EquipmentRequest to Pb message",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.InfoKV(ctx, describeEquipmentRequestV1LogTag, "success")

	return &pb.DescribeEquipmentRequestV1Response{
		EquipmentRequest: equipmentRequestPb,
	}, nil
}

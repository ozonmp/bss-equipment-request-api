package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) DescribeEquipmentRequestV1(
	ctx context.Context,
	req *pb.DescribeEquipmentRequestV1Request,
) (*pb.DescribeEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequest, err := o.equipmentRequestService.DescribeEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequest == nil {
		log.Debug().Uint64("equipmentRequestId", req.GetEquipmentRequestId()).Msg("equipment request not found")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	log.Debug().Msg("DescribeEquipmentRequestV1 - success")

	equipmentRequestPb, err := o.convertEquipmentRequestToPb(equipmentRequest)

	if err != nil {
		log.Error().Err(err).Msg("DescribeEquipmentRequestV1.convertEquipmentRequestToPb -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DescribeEquipmentRequestV1Response{
		EquipmentRequest: equipmentRequestPb,
	}, nil
}

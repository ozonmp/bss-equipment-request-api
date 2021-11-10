package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) ListEquipmentRequestV1(
	ctx context.Context,
	req *pb.ListEquipmentRequestV1Request,
) (*pb.ListEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	equipmentRequests, err := o.equipmentRequestService.ListEquipmentRequest(ctx, req.Limit, req.Offset)

	if err != nil {
		log.Error().Err(err).Msg("ListEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequests == nil {
		log.Debug().Msg(
			"unable to get list of equipment requests")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "unable to get list of equipment requests")
	}

	log.Debug().Msg("ListEquipmentRequestV1 - success")

	equipmentRequestPb, err := o.convertRepeatedEquipmentRequestsToPb(equipmentRequests)

	if err != nil {
		log.Error().Err(err).Msg("ListEquipmentRequestV1.convertRepeatedEquipmentRequestsToPb -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ListEquipmentRequestV1Response{
		Items: equipmentRequestPb,
	}, nil
}

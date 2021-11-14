package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) RemoveEquipmentRequestV1(
	ctx context.Context,
	req *pb.RemoveEquipmentRequestV1Request,
) (*pb.RemoveEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	exists, err := o.equipmentRequestService.CheckExistsEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !exists {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentRequestId).Msg("equipment request not found")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	result, err := o.equipmentRequestService.RemoveEquipmentRequest(ctx, req.EquipmentRequestId)

	if err != nil {
		log.Error().Err(err).Msg("RemoveEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if !result {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentRequestId).Msg("unable to remove equipment request")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "equipment request not removed")
	}

	log.Debug().Msg("RemoveEquipmentRequestV1 - success")

	return &pb.RemoveEquipmentRequestV1Response{
		Removed: result,
	}, nil
}

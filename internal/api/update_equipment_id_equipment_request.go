package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) UpdateEquipmentIDEquipmentRequestV1(
	ctx context.Context,
	req *pb.UpdateEquipmentIDEquipmentRequestV1Request,
) (*pb.UpdateEquipmentIDEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("UpdateEquipmentIdEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	exists, err := o.equipmentRequestService.CheckExistsEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		log.Error().Err(err).Msg("UpdateEquipmentIdEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if exists == false {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentId).Msg("equipment request not found")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	result, err := o.equipmentRequestService.UpdateEquipmentIDEquipmentRequest(ctx, req.EquipmentRequestId, req.EquipmentId)

	if err != nil {
		log.Error().Err(err).Msg("UpdateEquipmentIdEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if result == false {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentRequestId).Uint64(
			"equipmentId", req.EquipmentId).Msg("unable to update equipment id of equipment request")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "equipment id of equipment request not updated")
	}

	log.Debug().Msg("UpdateEquipmentIdEquipmentRequestV1 - success")

	return &pb.UpdateEquipmentIDEquipmentRequestV1Response{
		Updated: result,
	}, nil
}

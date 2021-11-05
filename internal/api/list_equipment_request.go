package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) ListEquipmentRequestV1(ctx context.Context, empty *emptypb.Empty) (
	*pb.ListEquipmentRequestV1Response, error) {

	equipmentRequests, err := o.equipmentRequestService.ListEquipmentRequest(ctx)

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

	equipmentRequestPb := o.convertRepeatedEquipmentRequestsToPb(equipmentRequests)

	return &pb.ListEquipmentRequestV1Response{
		Items: equipmentRequestPb,
	}, nil
}

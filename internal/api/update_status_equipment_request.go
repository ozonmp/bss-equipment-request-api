package api

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) UpdateStatusEquipmentRequestV1(
	ctx context.Context,
	req *pb.UpdateStatusEquipmentRequestV1Request,
) (*pb.UpdateStatusEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("UpdateStatusEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	exists, err := o.equipmentRequestService.CheckExistsEquipmentRequest(ctx, req.EquipmentRequestId)
	if err != nil {
		log.Error().Err(err).Msg("UpdateStatusEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if exists == false {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentRequestId).Int32(
			"equipmentRequestStatus", int32(req.EquipmentRequestStatus)).Msg("equipment request not found")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	statusVal, ok := pb.EquipmentRequestStatus_name[int32(req.EquipmentRequestStatus)]
	var equipmentRequestStatus model.EquipmentRequestStatus

	if ok {
		equipmentRequestStatus = model.EquipmentRequestStatus(statusVal)
	} else {
		return nil, ErrUnableToConvertEquipmentRequestStatus
	}

	result, err := o.equipmentRequestService.UpdateStatusEquipmentRequest(ctx, req.EquipmentRequestId, equipmentRequestStatus)

	if err != nil {
		log.Error().Err(err).Msg("UpdateStatusEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if result == false {
		log.Debug().Uint64("equipmentRequestId", req.EquipmentRequestId).Int32(
			"equipmentRequestStatus", int32(req.EquipmentRequestStatus)).Msg("unable to update update status of equipment request")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "status of equipment request equipment request not updated")
	}

	log.Debug().Msg("UpdateStatusEquipmentRequestV1 - success")

	return &pb.UpdateStatusEquipmentRequestV1Response{
		Updated: result,
	}, nil
}

package api

import (
	"context"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) CreateEquipmentRequestV1(
	ctx context.Context,
	req *pb.CreateEquipmentRequestV1Request,
) (*pb.CreateEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateEquipmentRequestV1Request - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := o.equipmentRequestService.CreateEquipmentRequest(ctx,
		req.GetEquipmentId(),
		req.GetEmployeeId(),
		req.GetCreatedAt(),
		req.GetDoneAt(),
		req.GetEquipmentRequestStatusId())

	if err != nil {
		log.Error().Err(err).Msg("CreateEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if id == 0 {
		log.Debug().Uint64(
			"employeeId", req.GetEmployeeId()).Uint64(
			"equipmentId", req.GetEquipmentId()).Time(
			"createdAt", req.GetCreatedAt().AsTime()).Time(
			"doneAt", req.GetDoneAt().AsTime()).Int32(
			"equipmentRequestStatusId", int32(req.GetEquipmentRequestStatusId())).Msg(
			"equipment request does not created")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "equipment request does not created")
	}

	log.Debug().Msg("CreateEquipmentRequestV1 - success")

	return &pb.CreateEquipmentRequestV1Response{
		EquipmentRequestId: id,
	}, nil
}

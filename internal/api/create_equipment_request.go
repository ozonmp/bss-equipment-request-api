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
		req.EquipmentId,
		req.EmployeeId,
		req.CreatedAt,
		req.DoneAt,
		req.EquipmentRequestStatusId)

	if err != nil {
		log.Error().Err(err).Msg("CreateEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if id == 0 {
		log.Debug().Uint64(
			"employeeId", req.EquipmentId).Uint64(
			"equipmentId", req.EmployeeId).Time(
			"createdAt", req.CreatedAt.AsTime()).Time(
			"doneAt", req.DoneAt.AsTime()).Int32(
			"equipmentRequestStatusId", int32(req.EquipmentRequestStatusId)).Msg(
			"equipment request does not created")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "equipment request does not created")
	}

	log.Debug().Msg("CreateEquipmentRequestV1 - success")

	return &pb.CreateEquipmentRequestV1Response{
		EquipmentRequestId: id,
	}, nil
}

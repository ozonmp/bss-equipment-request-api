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

	newItem := pb.EquipmentRequest{
		EquipmentId:            req.EquipmentId,
		EmployeeId:             req.EmployeeId,
		CreatedAt:              req.CreatedAt,
		UpdatedAt:              req.UpdatedAt,
		DeletedAt:              req.DeletedAt,
		DoneAt:                 req.DoneAt,
		EquipmentRequestStatus: req.EquipmentRequestStatus,
	}

	equipmentRequest, err := o.convertPbToEquipmentRequest(&newItem)

	if err != nil {
		log.Error().Err(err).Msg("CreateEquipmentRequestV1.convertPbToEquipmentRequest -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	id, err := o.equipmentRequestService.CreateEquipmentRequest(ctx, equipmentRequest)

	if err != nil {
		log.Error().Err(err).Msg("CreateEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if id == 0 {
		log.Debug().Uint64(
			"equipmentId", req.EquipmentId).Uint64(
			"employeeId", req.EmployeeId).Time(
			"createdAt", req.CreatedAt.AsTime()).Time(
			"updatedAt", req.UpdatedAt.AsTime()).Time(
			"deletedAt", req.DeletedAt.AsTime()).Time(
			"doneAt", req.DoneAt.AsTime()).Int32(
			"equipmentRequestStatus", int32(req.EquipmentRequestStatus)).Msg(
			"equipment request does not created")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.Internal, "equipment request does not created")
	}

	log.Debug().Msg("CreateEquipmentRequestV1 - success")

	return &pb.CreateEquipmentRequestV1Response{
		EquipmentRequestId: id,
	}, nil
}

package api

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *equipmentRequestAPI) CreateEquipmentRequestV1(
	ctx context.Context,
	req *pb.CreateEquipmentRequestV1Request,
) (*pb.CreateEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		logger.ErrorKV(ctx, createEquipmentRequestV1LogTag+": invalid argument",
			"err", err,
		)

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
		logger.FatalKV(ctx, createEquipmentRequestV1LogTag+": unable to convert Pb message to EquipmentRequest",
			"err", err,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	id, err := o.equipmentRequestService.CreateEquipmentRequest(ctx, equipmentRequest)

	if err != nil {
		logger.ErrorKV(ctx, createEquipmentRequestV1LogTag+": equipmentRequestService.CreateEquipmentRequest failed",
			"err", err,
			"equipmentId", req.EquipmentId,
			"employeeId", req.EmployeeId,
			"createdAt", req.CreatedAt,
			"updatedAt", req.UpdatedAt,
			"deletedAt", req.DeletedAt,
			"doneAt", req.DoneAt,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, err.Error())
	}

	if id == 0 {
		logger.DebugKV(ctx, createEquipmentRequestV1LogTag+": equipmentRequestService.CreateEquipmentRequest failed",
			"err", "unable to get created equipment request",
			"equipmentId", req.EquipmentId,
			"employeeId", req.EmployeeId,
			"createdAt", req.CreatedAt,
			"updatedAt", req.UpdatedAt,
			"deletedAt", req.DeletedAt,
			"doneAt", req.DoneAt,
			"equipmentRequestStatus", req.EquipmentRequestStatus,
		)

		return nil, status.Error(codes.Internal, "unable to get created equipment request")
	}

	logger.InfoKV(ctx, createEquipmentRequestV1LogTag, "success",
		"equipmentRequestId", id,
	)

	return &pb.CreateEquipmentRequestV1Response{
		EquipmentRequestId: id,
	}, nil
}

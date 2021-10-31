package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/bss-equipment-request-api/internal/repo"

	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
)

var (
	totalEquipmentRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_equipment_request_api_not_found_total",
		Help: "Total number of equipment requests that were not found",
	})
)

type equipmentRequestAPI struct {
	pb.UnimplementedBssEquipmentRequestApiServiceServer
	repo repo.Repo
}

// NewEquipmentRequestAPI returns api of bss-equipment-request-api service
func NewEquipmentRequestAPI(r repo.Repo) pb.BssEquipmentRequestApiServiceServer {
	return &equipmentRequestAPI{repo: r}
}

func (o *equipmentRequestAPI) DescribeEquipmentRequestV1(
	ctx context.Context,
	req *pb.DescribeEquipmentRequestV1Request,
) (*pb.DescribeEquipmentRequestV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeEquipmentRequestV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	equipmentRequest, err := o.repo.DescribeEquipmentRequest(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("DescribeEquipmentRequestV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if equipmentRequest == nil {
		log.Debug().Uint64("equipmentRequestId", req.Id).Msg("equipment request not found")
		totalEquipmentRequestNotFound.Inc()

		return nil, status.Error(codes.NotFound, "equipment request not found")
	}

	log.Debug().Msg("DescribeEquipmentRequestV1 - success")

	return &pb.DescribeEquipmentRequestV1Response{
		Result: &pb.EquipmentRequest{
			Id:  equipmentRequest.Id,
			EmployeeId: equipmentRequest.EmployeeId,
			EquipmentId: equipmentRequest.EquipmentId,
			CreatedAt: timestamppb.New(equipmentRequest.CreatedAt),
			DoneAt: timestamppb.New(equipmentRequest.DoneAt),
			EquipmentRequestStatusId: equipmentRequest.EquipmentRequestStatusId,
		},
	}, nil
}

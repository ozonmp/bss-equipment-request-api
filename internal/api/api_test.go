package api_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/api"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks/server"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	pb "github.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

var listener *bufconn.Listener

func dbMock() (sqlmock.Sqlmock, *sqlx.DB) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	return mock, sqlxDB
}

func setUp(t *testing.T) (*grpc.ClientConn, context.Context, func(*grpc.ClientConn)) {
	listener = bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()

	ctrl := gomock.NewController(t)

	_, sqlxDB := dbMock()
	repo := mocks.NewMockEquipmentRequestRepo(ctrl)
	eventRepo := mocks.NewMockEventRepo(ctrl)

	equipmentRequestService := equipment_request.New(sqlxDB, repo, eventRepo)

	pb.RegisterBssEquipmentRequestApiServiceServer(grpcServer, api.NewEquipmentRequestAPI(equipmentRequestService))

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithInsecure(), grpc.WithContextDialer(bufDialer))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	closeFunc := func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Unable to close connection")
		}
	}

	return conn, ctx, closeFunc
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}

func Test_DescribeEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.DescribeEquipmentRequestV1Request{}

	equipmentRequest, err := client.DescribeEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_DescribeEquipmentRequestV1_WrongFormat(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.DescribeEquipmentRequestV1Request{
		EquipmentRequestId: 0,
	}

	equipmentRequest, err := client.DescribeEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())

}

func Test_RemoveEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.RemoveEquipmentRequestV1Request{}

	equipmentRequest, err := client.RemoveEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_RemoveEquipmentRequestV1_WrongFormat(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.RemoveEquipmentRequestV1Request{
		EquipmentRequestId: 0,
	}

	equipmentRequest, err := client.RemoveEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEmployeeId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:             0,
		EquipmentId:            12,
		EquipmentRequestStatus: 2,
		CreatedAt:              timestamppb.New(time.Now()),
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEquipmentId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:             2,
		EquipmentId:            0,
		EquipmentRequestStatus: 0,
		CreatedAt:              timestamppb.New(time.Now()),
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEquipmentRequestStatusId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:             2,
		EquipmentId:            1,
		EquipmentRequestStatus: 10,
		CreatedAt:              timestamppb.New(time.Now()),
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongCreatedAt(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:             2,
		EquipmentId:            1,
		EquipmentRequestStatus: 10,
		CreatedAt:              nil,
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateEquipmentIdEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateEquipmentIDEquipmentRequestV1Request{}

	equipmentRequest, err := client.UpdateEquipmentIDEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateEquipmentIdEquipmentRequestV1_WrongEquipmentId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateEquipmentIDEquipmentRequestV1Request{
		EquipmentId:        0,
		EquipmentRequestId: 10,
	}

	equipmentRequest, err := client.UpdateEquipmentIDEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateEquipmentIdEquipmentRequestV1_WrongEquipmentRequestId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateEquipmentIDEquipmentRequestV1Request{
		EquipmentId:        22,
		EquipmentRequestId: 0,
	}

	equipmentRequest, err := client.UpdateEquipmentIDEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateStatusEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateStatusEquipmentRequestV1Request{}

	equipmentRequest, err := client.UpdateStatusEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateStatusEquipmentRequestV1_WrongStatus(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateStatusEquipmentRequestV1Request{
		EquipmentRequestStatus: 22,
		EquipmentRequestId:     10,
	}

	equipmentRequest, err := client.UpdateStatusEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_UpdateStatusEquipmentRequestV1_WrongEquipmentRequestId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.UpdateStatusEquipmentRequestV1Request{
		EquipmentRequestStatus: 7,
		EquipmentRequestId:     0,
	}

	equipmentRequest, err := client.UpdateStatusEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_ListEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.ListEquipmentRequestV1Request{}

	equipmentRequest, err := client.ListEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_ListEquipmentRequestV1_WrongLimit(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.ListEquipmentRequestV1Request{
		Limit:  0,
		Offset: 10,
	}

	equipmentRequest, err := client.ListEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_ListEquipmentRequestV1_WrongPerPage(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.ListEquipmentRequestV1Request{
		Limit:  7,
		Offset: 34,
	}

	equipmentRequest, err := client.ListEquipmentRequestV1(ctx, &request)

	require.NotNil(t, err)
	require.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	require.Equal(t, codes.InvalidArgument, er.Code())
}

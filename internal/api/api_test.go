package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/api"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks/server"
	"github.com/stretchr/testify/assert"
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

func setUp(t *testing.T) (*grpc.ClientConn, context.Context, func(*grpc.ClientConn)) {
	listener = bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockRepo(ctrl)
	equipmentRequestService := equipment_request.New(repo)

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

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_DescribeEquipmentRequestV1_WrongFormat(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.DescribeEquipmentRequestV1Request{
		EquipmentRequestId: 0,
	}

	equipmentRequest, err := client.DescribeEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())

}

func Test_RemoveEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.RemoveEquipmentRequestV1Request{}

	equipmentRequest, err := client.RemoveEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_RemoveEquipmentRequestV1_WrongFormat(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.RemoveEquipmentRequestV1Request{
		EquipmentRequestId: 0,
	}

	equipmentRequest, err := client.RemoveEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_EmptyRequest(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEmployeeId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:               0,
		EquipmentId:              12,
		EquipmentRequestStatusId: 2,
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEquipmentId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:               2,
		EquipmentId:              0,
		EquipmentRequestStatusId: 0,
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

func Test_CreateEquipmentRequestV1_WrongEquipmentRequestStatusId(t *testing.T) {
	conn, ctx, closeFunc := setUp(t)

	defer closeFunc(conn)

	client := pb.NewBssEquipmentRequestApiServiceClient(conn)

	request := pb.CreateEquipmentRequestV1Request{
		EmployeeId:               2,
		EquipmentId:              1,
		EquipmentRequestStatusId: 10,
	}

	equipmentRequest, err := client.CreateEquipmentRequestV1(ctx, &request)

	assert.NotNil(t, err)
	assert.Nil(t, equipmentRequest)

	er, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, er.Code())
}

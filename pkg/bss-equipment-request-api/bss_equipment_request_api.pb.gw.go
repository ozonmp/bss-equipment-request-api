// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ozonmp/bss_equipment_request_api/v1/bss_equipment_request_api.proto

/*
Package bss_equipment_request_api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package bss_equipment_request_api

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DescribeEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["equipment_request_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "equipment_request_id")
	}

	protoReq.EquipmentRequestId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "equipment_request_id", err)
	}

	msg, err := client.DescribeEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DescribeEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["equipment_request_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "equipment_request_id")
	}

	protoReq.EquipmentRequestId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "equipment_request_id", err)
	}

	msg, err := server.DescribeEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

func request_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

func request_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ListEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

func request_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RemoveEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.RemoveEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RemoveEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.RemoveEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

func request_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateEquipmentIdEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateEquipmentIdEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateEquipmentIdEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateEquipmentIdEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

func request_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, client BssEquipmentRequestApiServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateStatusEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UpdateStatusEquipmentRequestV1(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(ctx context.Context, marshaler runtime.Marshaler, server BssEquipmentRequestApiServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateStatusEquipmentRequestV1Request
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UpdateStatusEquipmentRequestV1(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBssEquipmentRequestApiServiceHandlerServer registers the http handlers for service BssEquipmentRequestApiService to "mux".
// UnaryRPC     :call BssEquipmentRequestApiServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBssEquipmentRequestApiServiceHandlerFromEndpoint instead.
func RegisterBssEquipmentRequestApiServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BssEquipmentRequestApiServiceServer) error {

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/{equipment_request_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/create"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_ListEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/list"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/remove"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateEquipmentIdEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/update/equipment_id"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateStatusEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/update/status"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBssEquipmentRequestApiServiceHandlerFromEndpoint is same as RegisterBssEquipmentRequestApiServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBssEquipmentRequestApiServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBssEquipmentRequestApiServiceHandler(ctx, mux, conn)
}

// RegisterBssEquipmentRequestApiServiceHandler registers the http handlers for service BssEquipmentRequestApiService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBssEquipmentRequestApiServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBssEquipmentRequestApiServiceHandlerClient(ctx, mux, NewBssEquipmentRequestApiServiceClient(conn))
}

// RegisterBssEquipmentRequestApiServiceHandlerClient registers the http handlers for service BssEquipmentRequestApiService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BssEquipmentRequestApiServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BssEquipmentRequestApiServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BssEquipmentRequestApiServiceClient" to call the correct interceptors.
func RegisterBssEquipmentRequestApiServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BssEquipmentRequestApiServiceClient) error {

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/{equipment_request_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/create"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_ListEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/list"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_ListEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/equipment_requests/remove"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateEquipmentIdEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/update/equipment_id"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateStatusEquipmentRequestV1", runtime.WithHTTPPathPattern("/api/v1/update/status"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v1", "equipment_requests", "equipment_request_id"}, ""))

	pattern_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "equipment_requests", "create"}, ""))

	pattern_BssEquipmentRequestApiService_ListEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "equipment_requests", "list"}, ""))

	pattern_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "equipment_requests", "remove"}, ""))

	pattern_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "update", "equipment_id"}, ""))

	pattern_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "update", "status"}, ""))
)

var (
	forward_BssEquipmentRequestApiService_DescribeEquipmentRequestV1_0 = runtime.ForwardResponseMessage

	forward_BssEquipmentRequestApiService_CreateEquipmentRequestV1_0 = runtime.ForwardResponseMessage

	forward_BssEquipmentRequestApiService_ListEquipmentRequestV1_0 = runtime.ForwardResponseMessage

	forward_BssEquipmentRequestApiService_RemoveEquipmentRequestV1_0 = runtime.ForwardResponseMessage

	forward_BssEquipmentRequestApiService_UpdateEquipmentIdEquipmentRequestV1_0 = runtime.ForwardResponseMessage

	forward_BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_0 = runtime.ForwardResponseMessage
)

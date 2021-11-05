# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/bss_equipment_request_api/v1/bss_equipment_request_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2


class BssEquipmentRequestApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribeEquipmentRequestV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Request, ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def CreateEquipmentRequestV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Request, ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListEquipmentRequestV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Request, ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemoveEquipmentRequestV1(self, stream: 'grpclib.server.Stream[ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Request, ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1': grpclib.const.Handler(
                self.DescribeEquipmentRequestV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Request,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Response,
            ),
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1': grpclib.const.Handler(
                self.CreateEquipmentRequestV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Request,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Response,
            ),
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1': grpclib.const.Handler(
                self.ListEquipmentRequestV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Request,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Response,
            ),
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1': grpclib.const.Handler(
                self.RemoveEquipmentRequestV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Request,
                ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Response,
            ),
        }


class BssEquipmentRequestApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribeEquipmentRequestV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1',
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Request,
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.DescribeEquipmentRequestV1Response,
        )
        self.CreateEquipmentRequestV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1',
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Request,
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.CreateEquipmentRequestV1Response,
        )
        self.ListEquipmentRequestV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1',
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Request,
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.ListEquipmentRequestV1Response,
        )
        self.RemoveEquipmentRequestV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1',
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Request,
            ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2.RemoveEquipmentRequestV1Response,
        )

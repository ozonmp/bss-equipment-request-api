import asyncio

from grpclib.client import Channel

from ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_grpc import BssEquipmentRequestApiServiceStub
from ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2 import DescribeEquipmentRequestV1

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = BssEquipmentRequestApiServiceStub(channel)

        req = DescribeEquipmentRequestV1Request(equipment_request_id=1)
        reply = await client.DescribeEquipmentRequestV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())

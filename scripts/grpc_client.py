import asyncio

from grpclib.client import Channel

from ozonmp.omp_template_api.v1.omp_template_api_grpc import OmpTemplateApiServiceStub
from ozonmp.omp_template_api.v1.omp_template_api_pb2 import DescribeTemplateV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = OmpTemplateApiServiceStub(channel)

        req = DescribeTemplateV1Request(id=1)
        reply = await client.DescribeTemplateV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())

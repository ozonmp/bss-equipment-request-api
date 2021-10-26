import grpc
import ozonmp.omp_template_api.v1.omp_template_api_pb2_grpc as pb2_grpc
import ozonmp.omp_template_api.v1.omp_template_api_pb2 as pb2

template_api_addr = '0.0.0.0:8082'
with grpc.insecure_channel(template_api_addr) as channel:
    template_api_client = pb2_grpc.OmpTemplateApiServiceStub(channel)
    meta = (('omp-source', 'jupyterhub'),)
    req = pb2.DescribeTemplateV1Request(id=1)
    resp = template_api_client.DescribeTemplateV1(request=req, metadata=meta)
    print(resp.value)
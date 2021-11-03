#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1"

payload=$(
  cat <<EOF
{

}
EOF
)

grpcurl -plaintext -emit-defaults \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1"

payload=$(
  cat <<EOF
{
  "equipment_request_id": 70011
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
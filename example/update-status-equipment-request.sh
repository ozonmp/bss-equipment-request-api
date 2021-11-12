#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateStatusEquipmentRequestV1"

payload=$(
  cat <<EOF
{
  "equipment_request_id": 70010,
  "equipment_request_status": 3
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
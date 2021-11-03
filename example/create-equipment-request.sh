#!/bin/sh

GRPC_HOST="localhost:8082"
GRPC_METHOD="ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1"

payload=$(
  cat <<EOF
{
  "employee_id": 12,
  "equipment_id": 12,
  "created_at": "2020-05-22T20:32:05Z",
  "done_at": "2020-05-22T23:32:05Z",
  "equipment_request_status_id": 0
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
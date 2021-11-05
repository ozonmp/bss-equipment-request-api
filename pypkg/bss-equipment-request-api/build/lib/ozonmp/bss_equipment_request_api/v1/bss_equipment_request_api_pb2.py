# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/bss_equipment_request_api/v1/bss_equipment_request_api.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\nCozonmp/bss_equipment_request_api/v1/bss_equipment_request_api.proto\x12#ozonmp.bss_equipment_request_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x95\x02\n\x10\x45quipmentRequest\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x1f\n\x0b\x65mployee_id\x18\x02 \x01(\x04R\nemployeeId\x12!\n\x0c\x65quipment_id\x18\x03 \x01(\x04R\x0b\x65quipmentId\x12\x39\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x33\n\x07\x64one_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x06\x64oneAt\x12=\n\x1b\x65quipment_request_status_id\x18\x06 \x01(\x04R\x18\x65quipmentRequestStatusId\"^\n!DescribeEquipmentRequestV1Request\x12\x39\n\x14\x65quipment_request_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x12\x65quipmentRequestId\"\x88\x01\n\"DescribeEquipmentRequestV1Response\x12\x62\n\x11\x65quipment_request\x18\x01 \x01(\x0b\x32\x35.ozonmp.bss_equipment_request_api.v1.EquipmentRequestR\x10\x65quipmentRequest\"\xee\x02\n\x1f\x43reateEquipmentRequestV1Request\x12(\n\x0b\x65mployee_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\nemployeeId\x12*\n\x0c\x65quipment_id\x18\x02 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0b\x65quipmentId\x12\x39\n\ncreated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x33\n\x07\x64one_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x06\x64oneAt\x12\x84\x01\n\x1b\x65quipment_request_status_id\x18\x05 \x01(\x0e\x32;.ozonmp.bss_equipment_request_api.v1.EquipmentRequestStatusB\x08\xfa\x42\x05\x82\x01\x02\x10\x01R\x18\x65quipmentRequestStatusId\"T\n CreateEquipmentRequestV1Response\x12\x30\n\x14\x65quipment_request_id\x18\x01 \x01(\x04R\x12\x65quipmentRequestId\"\x1f\n\x1dListEquipmentRequestV1Request\"m\n\x1eListEquipmentRequestV1Response\x12K\n\x05items\x18\x01 \x03(\x0b\x32\x35.ozonmp.bss_equipment_request_api.v1.EquipmentRequestR\x05items\"\\\n\x1fRemoveEquipmentRequestV1Request\x12\x39\n\x14\x65quipment_request_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x12\x65quipmentRequestId\"<\n RemoveEquipmentRequestV1Response\x12\x18\n\x07removed\x18\x01 \x01(\x08R\x07removed*\xe7\x01\n\x16\x45quipmentRequestStatus\x12+\n\'EQUIPMENT_REQUEST_STATUS_ID_UNSPECIFIED\x10\x00\x12\"\n\x1e\x45QUIPMENT_REQUEST_STATUS_ID_DO\x10\x01\x12+\n\'EQUIPMENT_REQUEST_STATUS_ID_IN_PROGRESS\x10\x02\x12$\n EQUIPMENT_REQUEST_STATUS_ID_DONE\x10\x03\x12)\n%EQUIPMENT_REQUEST_STATUS_ID_CANCELLED\x10\x04\x32\x8d\x07\n\x1d\x42ssEquipmentRequestApiService\x12\xeb\x01\n\x1a\x44\x65scribeEquipmentRequestV1\x12\x46.ozonmp.bss_equipment_request_api.v1.DescribeEquipmentRequestV1Request\x1aG.ozonmp.bss_equipment_request_api.v1.DescribeEquipmentRequestV1Response\"<\x82\xd3\xe4\x93\x02\x36\"1/api/v1/equipment_requests/{equipment_request_id}:\x01*\x12\xd5\x01\n\x18\x43reateEquipmentRequestV1\x12\x44.ozonmp.bss_equipment_request_api.v1.CreateEquipmentRequestV1Request\x1a\x45.ozonmp.bss_equipment_request_api.v1.CreateEquipmentRequestV1Response\",\x82\xd3\xe4\x93\x02&\"!/api/v1/equipment_requests/create:\x01*\x12\xcd\x01\n\x16ListEquipmentRequestV1\x12\x42.ozonmp.bss_equipment_request_api.v1.ListEquipmentRequestV1Request\x1a\x43.ozonmp.bss_equipment_request_api.v1.ListEquipmentRequestV1Response\"*\x82\xd3\xe4\x93\x02$\"\x1f/api/v1/equipment_requests/list:\x01*\x12\xd5\x01\n\x18RemoveEquipmentRequestV1\x12\x44.ozonmp.bss_equipment_request_api.v1.RemoveEquipmentRequestV1Request\x1a\x45.ozonmp.bss_equipment_request_api.v1.RemoveEquipmentRequestV1Response\",\x82\xd3\xe4\x93\x02&\"!/api/v1/equipment_requests/remove:\x01*BeZcgithub.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api;bss_equipment_request_apib\x06proto3')

_EQUIPMENTREQUESTSTATUS = DESCRIPTOR.enum_types_by_name['EquipmentRequestStatus']
EquipmentRequestStatus = enum_type_wrapper.EnumTypeWrapper(_EQUIPMENTREQUESTSTATUS)
EQUIPMENT_REQUEST_STATUS_ID_UNSPECIFIED = 0
EQUIPMENT_REQUEST_STATUS_ID_DO = 1
EQUIPMENT_REQUEST_STATUS_ID_IN_PROGRESS = 2
EQUIPMENT_REQUEST_STATUS_ID_DONE = 3
EQUIPMENT_REQUEST_STATUS_ID_CANCELLED = 4


_EQUIPMENTREQUEST = DESCRIPTOR.message_types_by_name['EquipmentRequest']
_DESCRIBEEQUIPMENTREQUESTV1REQUEST = DESCRIPTOR.message_types_by_name['DescribeEquipmentRequestV1Request']
_DESCRIBEEQUIPMENTREQUESTV1RESPONSE = DESCRIPTOR.message_types_by_name['DescribeEquipmentRequestV1Response']
_CREATEEQUIPMENTREQUESTV1REQUEST = DESCRIPTOR.message_types_by_name['CreateEquipmentRequestV1Request']
_CREATEEQUIPMENTREQUESTV1RESPONSE = DESCRIPTOR.message_types_by_name['CreateEquipmentRequestV1Response']
_LISTEQUIPMENTREQUESTV1REQUEST = DESCRIPTOR.message_types_by_name['ListEquipmentRequestV1Request']
_LISTEQUIPMENTREQUESTV1RESPONSE = DESCRIPTOR.message_types_by_name['ListEquipmentRequestV1Response']
_REMOVEEQUIPMENTREQUESTV1REQUEST = DESCRIPTOR.message_types_by_name['RemoveEquipmentRequestV1Request']
_REMOVEEQUIPMENTREQUESTV1RESPONSE = DESCRIPTOR.message_types_by_name['RemoveEquipmentRequestV1Response']
EquipmentRequest = _reflection.GeneratedProtocolMessageType('EquipmentRequest', (_message.Message,), {
  'DESCRIPTOR' : _EQUIPMENTREQUEST,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.EquipmentRequest)
  })
_sym_db.RegisterMessage(EquipmentRequest)

DescribeEquipmentRequestV1Request = _reflection.GeneratedProtocolMessageType('DescribeEquipmentRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEEQUIPMENTREQUESTV1REQUEST,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.DescribeEquipmentRequestV1Request)
  })
_sym_db.RegisterMessage(DescribeEquipmentRequestV1Request)

DescribeEquipmentRequestV1Response = _reflection.GeneratedProtocolMessageType('DescribeEquipmentRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEEQUIPMENTREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.DescribeEquipmentRequestV1Response)
  })
_sym_db.RegisterMessage(DescribeEquipmentRequestV1Response)

CreateEquipmentRequestV1Request = _reflection.GeneratedProtocolMessageType('CreateEquipmentRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _CREATEEQUIPMENTREQUESTV1REQUEST,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.CreateEquipmentRequestV1Request)
  })
_sym_db.RegisterMessage(CreateEquipmentRequestV1Request)

CreateEquipmentRequestV1Response = _reflection.GeneratedProtocolMessageType('CreateEquipmentRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _CREATEEQUIPMENTREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.CreateEquipmentRequestV1Response)
  })
_sym_db.RegisterMessage(CreateEquipmentRequestV1Response)

ListEquipmentRequestV1Request = _reflection.GeneratedProtocolMessageType('ListEquipmentRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _LISTEQUIPMENTREQUESTV1REQUEST,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.ListEquipmentRequestV1Request)
  })
_sym_db.RegisterMessage(ListEquipmentRequestV1Request)

ListEquipmentRequestV1Response = _reflection.GeneratedProtocolMessageType('ListEquipmentRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _LISTEQUIPMENTREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.ListEquipmentRequestV1Response)
  })
_sym_db.RegisterMessage(ListEquipmentRequestV1Response)

RemoveEquipmentRequestV1Request = _reflection.GeneratedProtocolMessageType('RemoveEquipmentRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEEQUIPMENTREQUESTV1REQUEST,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.RemoveEquipmentRequestV1Request)
  })
_sym_db.RegisterMessage(RemoveEquipmentRequestV1Request)

RemoveEquipmentRequestV1Response = _reflection.GeneratedProtocolMessageType('RemoveEquipmentRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEEQUIPMENTREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.bss_equipment_request_api.v1.bss_equipment_request_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_equipment_request_api.v1.RemoveEquipmentRequestV1Response)
  })
_sym_db.RegisterMessage(RemoveEquipmentRequestV1Response)

_BSSEQUIPMENTREQUESTAPISERVICE = DESCRIPTOR.services_by_name['BssEquipmentRequestApiService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Zcgithub.com/ozonmp/bss-equipment-request-api/pkg/bss-equipment-request-api;bss_equipment_request_api'
  _DESCRIBEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_id']._options = None
  _DESCRIBEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_id']._serialized_options = b'\372B\0042\002 \000'
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['employee_id']._options = None
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['employee_id']._serialized_options = b'\372B\0042\002 \000'
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_id']._options = None
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_id']._serialized_options = b'\372B\0042\002 \000'
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_status_id']._options = None
  _CREATEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_status_id']._serialized_options = b'\372B\005\202\001\002\020\001'
  _REMOVEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_id']._options = None
  _REMOVEEQUIPMENTREQUESTV1REQUEST.fields_by_name['equipment_request_id']._serialized_options = b'\372B\0042\002 \000'
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['DescribeEquipmentRequestV1']._options = None
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['DescribeEquipmentRequestV1']._serialized_options = b'\202\323\344\223\0026\"1/api/v1/equipment_requests/{equipment_request_id}:\001*'
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['CreateEquipmentRequestV1']._options = None
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['CreateEquipmentRequestV1']._serialized_options = b'\202\323\344\223\002&\"!/api/v1/equipment_requests/create:\001*'
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['ListEquipmentRequestV1']._options = None
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['ListEquipmentRequestV1']._serialized_options = b'\202\323\344\223\002$\"\037/api/v1/equipment_requests/list:\001*'
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['RemoveEquipmentRequestV1']._options = None
  _BSSEQUIPMENTREQUESTAPISERVICE.methods_by_name['RemoveEquipmentRequestV1']._serialized_options = b'\202\323\344\223\002&\"!/api/v1/equipment_requests/remove:\001*'
  _EQUIPMENTREQUESTSTATUS._serialized_start=1467
  _EQUIPMENTREQUESTSTATUS._serialized_end=1698
  _EQUIPMENTREQUEST._serialized_start=197
  _EQUIPMENTREQUEST._serialized_end=474
  _DESCRIBEEQUIPMENTREQUESTV1REQUEST._serialized_start=476
  _DESCRIBEEQUIPMENTREQUESTV1REQUEST._serialized_end=570
  _DESCRIBEEQUIPMENTREQUESTV1RESPONSE._serialized_start=573
  _DESCRIBEEQUIPMENTREQUESTV1RESPONSE._serialized_end=709
  _CREATEEQUIPMENTREQUESTV1REQUEST._serialized_start=712
  _CREATEEQUIPMENTREQUESTV1REQUEST._serialized_end=1078
  _CREATEEQUIPMENTREQUESTV1RESPONSE._serialized_start=1080
  _CREATEEQUIPMENTREQUESTV1RESPONSE._serialized_end=1164
  _LISTEQUIPMENTREQUESTV1REQUEST._serialized_start=1166
  _LISTEQUIPMENTREQUESTV1REQUEST._serialized_end=1197
  _LISTEQUIPMENTREQUESTV1RESPONSE._serialized_start=1199
  _LISTEQUIPMENTREQUESTV1RESPONSE._serialized_end=1308
  _REMOVEEQUIPMENTREQUESTV1REQUEST._serialized_start=1310
  _REMOVEEQUIPMENTREQUESTV1REQUEST._serialized_end=1402
  _REMOVEEQUIPMENTREQUESTV1RESPONSE._serialized_start=1404
  _REMOVEEQUIPMENTREQUESTV1RESPONSE._serialized_end=1464
  _BSSEQUIPMENTREQUESTAPISERVICE._serialized_start=1701
  _BSSEQUIPMENTREQUESTAPISERVICE._serialized_end=2610
# @@protoc_insertion_point(module_scope)
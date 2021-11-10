# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: google/api/http.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15google/api/http.proto\x12\ngoogle.api\"y\n\x04Http\x12*\n\x05rules\x18\x01 \x03(\x0b\x32\x14.google.api.HttpRuleR\x05rules\x12\x45\n\x1f\x66ully_decode_reserved_expansion\x18\x02 \x01(\x08R\x1c\x66ullyDecodeReservedExpansion\"\xda\x02\n\x08HttpRule\x12\x1a\n\x08selector\x18\x01 \x01(\tR\x08selector\x12\x12\n\x03get\x18\x02 \x01(\tH\x00R\x03get\x12\x12\n\x03put\x18\x03 \x01(\tH\x00R\x03put\x12\x14\n\x04post\x18\x04 \x01(\tH\x00R\x04post\x12\x18\n\x06\x64\x65lete\x18\x05 \x01(\tH\x00R\x06\x64\x65lete\x12\x16\n\x05patch\x18\x06 \x01(\tH\x00R\x05patch\x12\x37\n\x06\x63ustom\x18\x08 \x01(\x0b\x32\x1d.google.api.CustomHttpPatternH\x00R\x06\x63ustom\x12\x12\n\x04\x62ody\x18\x07 \x01(\tR\x04\x62ody\x12#\n\rresponse_body\x18\x0c \x01(\tR\x0cresponseBody\x12\x45\n\x13\x61\x64\x64itional_bindings\x18\x0b \x03(\x0b\x32\x14.google.api.HttpRuleR\x12\x61\x64\x64itionalBindingsB\t\n\x07pattern\";\n\x11\x43ustomHttpPattern\x12\x12\n\x04kind\x18\x01 \x01(\tR\x04kind\x12\x12\n\x04path\x18\x02 \x01(\tR\x04pathBj\n\x0e\x63om.google.apiB\tHttpProtoP\x01ZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations\xf8\x01\x01\xa2\x02\x04GAPIb\x06proto3')



_HTTP = DESCRIPTOR.message_types_by_name['Http']
_HTTPRULE = DESCRIPTOR.message_types_by_name['HttpRule']
_CUSTOMHTTPPATTERN = DESCRIPTOR.message_types_by_name['CustomHttpPattern']
Http = _reflection.GeneratedProtocolMessageType('Http', (_message.Message,), {
  'DESCRIPTOR' : _HTTP,
  '__module__' : 'google.api.http_pb2'
  # @@protoc_insertion_point(class_scope:google.api.Http)
  })
_sym_db.RegisterMessage(Http)

HttpRule = _reflection.GeneratedProtocolMessageType('HttpRule', (_message.Message,), {
  'DESCRIPTOR' : _HTTPRULE,
  '__module__' : 'google.api.http_pb2'
  # @@protoc_insertion_point(class_scope:google.api.HttpRule)
  })
_sym_db.RegisterMessage(HttpRule)

CustomHttpPattern = _reflection.GeneratedProtocolMessageType('CustomHttpPattern', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMHTTPPATTERN,
  '__module__' : 'google.api.http_pb2'
  # @@protoc_insertion_point(class_scope:google.api.CustomHttpPattern)
  })
_sym_db.RegisterMessage(CustomHttpPattern)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\016com.google.apiB\tHttpProtoP\001ZAgoogle.golang.org/genproto/googleapis/api/annotations;annotations\370\001\001\242\002\004GAPI'
  _HTTP._serialized_start=37
  _HTTP._serialized_end=158
  _HTTPRULE._serialized_start=161
  _HTTPRULE._serialized_end=507
  _CUSTOMHTTPPATTERN._serialized_start=509
  _CUSTOMHTTPPATTERN._serialized_end=568
# @@protoc_insertion_point(module_scope)

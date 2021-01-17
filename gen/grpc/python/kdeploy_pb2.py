# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kdeploy.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from github.com.mwitkow.go_proto_validators import validator_pb2 as github_dot_com_dot_mwitkow_dot_go__proto__validators_dot_validator__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='kdeploy.proto',
  package='kdeploy',
  syntax='proto3',
  serialized_options=_b('Z\tkdeploypb'),
  serialized_pb=_b('\n\rkdeploy.proto\x12\x07kdeploy\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19google/protobuf/any.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x36github.com/mwitkow/go-proto-validators/validator.proto\"\x8e\x02\n\x03\x41pp\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tnamespace\x18\x02 \x01(\t\x12\r\n\x05image\x18\x03 \x01(\t\x12\"\n\x03\x65nv\x18\x04 \x03(\x0b\x32\x15.kdeploy.App.EnvEntry\x12&\n\x05ports\x18\x05 \x03(\x0b\x32\x17.kdeploy.App.PortsEntry\x12\x10\n\x08replicas\x18\x06 \x01(\r\x12\x1f\n\x06status\x18\x07 \x01(\x0b\x32\x0f.kdeploy.Status\x1a*\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a,\n\nPortsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\"\xc4\x02\n\x0e\x41ppConstructor\x12\x1e\n\x04name\x18\x01 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12#\n\tnamespace\x18\x02 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12\x1f\n\x05image\x18\x03 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12-\n\x03\x65nv\x18\x04 \x03(\x0b\x32 .kdeploy.AppConstructor.EnvEntry\x12\x31\n\x05ports\x18\x05 \x03(\x0b\x32\".kdeploy.AppConstructor.PortsEntry\x12\x10\n\x08replicas\x18\x06 \x01(\r\x1a*\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a,\n\nPortsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\"\xa3\x02\n\tAppUpdate\x12\x1e\n\x04name\x18\x01 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12#\n\tnamespace\x18\x02 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12\r\n\x05image\x18\x03 \x01(\t\x12(\n\x03\x65nv\x18\x04 \x03(\x0b\x32\x1b.kdeploy.AppUpdate.EnvEntry\x12,\n\x05ports\x18\x05 \x03(\x0b\x32\x1d.kdeploy.AppUpdate.PortsEntry\x12\x10\n\x08replicas\x18\x06 \x01(\r\x1a*\n\x08\x45nvEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a,\n\nPortsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\"M\n\x06\x41ppRef\x12\x1e\n\x04name\x18\x01 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\x12#\n\tnamespace\x18\x02 \x01(\tB\x10\xe2\xdf\x1f\x0c\n\n^.{1,225}$\";\n\x07Replica\x12\r\n\x05phase\x18\x01 \x01(\t\x12\x11\n\tcondition\x18\x02 \x01(\t\x12\x0e\n\x06reason\x18\x03 \x01(\t\",\n\x06Status\x12\"\n\x08replicas\x18\x01 \x03(\x0b\x32\x10.kdeploy.Replica\"\x16\n\x03Log\x12\x0f\n\x07message\x18\x01 \x01(\t\"*\n\x04\x41pps\x12\"\n\x0c\x61pplications\x18\x01 \x03(\x0b\x32\x0c.kdeploy.App\"\x1e\n\tNamespace\x12\x11\n\tnamespace\x18\x01 \x01(\t\" \n\nNamespaces\x12\x12\n\nnamespaces\x18\x01 \x03(\t2\xf7\x02\n\x0eKdeployService\x12\x34\n\tCreateApp\x12\x17.kdeploy.AppConstructor\x1a\x0c.kdeploy.App\"\x00\x12/\n\tUpdateApp\x12\x12.kdeploy.AppUpdate\x1a\x0c.kdeploy.App\"\x00\x12\x36\n\tDeleteApp\x12\x0f.kdeploy.AppRef\x1a\x16.google.protobuf.Empty\"\x00\x12)\n\x06GetApp\x12\x0f.kdeploy.AppRef\x1a\x0c.kdeploy.App\"\x00\x12)\n\x04Logs\x12\x0f.kdeploy.AppRef\x1a\x0c.kdeploy.Log\"\x00\x30\x01\x12?\n\x0eListNamespaces\x12\x16.google.protobuf.Empty\x1a\x13.kdeploy.Namespaces\"\x00\x12/\n\x08ListApps\x12\x12.kdeploy.Namespace\x1a\r.kdeploy.Apps\"\x00\x42\x0bZ\tkdeploypbb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_any__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,github_dot_com_dot_mwitkow_dot_go__proto__validators_dot_validator__pb2.DESCRIPTOR,])




_APP_ENVENTRY = _descriptor.Descriptor(
  name='EnvEntry',
  full_name='kdeploy.App.EnvEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.App.EnvEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.App.EnvEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=384,
  serialized_end=426,
)

_APP_PORTSENTRY = _descriptor.Descriptor(
  name='PortsEntry',
  full_name='kdeploy.App.PortsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.App.PortsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.App.PortsEntry.value', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=428,
  serialized_end=472,
)

_APP = _descriptor.Descriptor(
  name='App',
  full_name='kdeploy.App',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='kdeploy.App.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='kdeploy.App.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image', full_name='kdeploy.App.image', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='env', full_name='kdeploy.App.env', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='kdeploy.App.ports', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='kdeploy.App.replicas', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='kdeploy.App.status', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_APP_ENVENTRY, _APP_PORTSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=202,
  serialized_end=472,
)


_APPCONSTRUCTOR_ENVENTRY = _descriptor.Descriptor(
  name='EnvEntry',
  full_name='kdeploy.AppConstructor.EnvEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.AppConstructor.EnvEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.AppConstructor.EnvEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=384,
  serialized_end=426,
)

_APPCONSTRUCTOR_PORTSENTRY = _descriptor.Descriptor(
  name='PortsEntry',
  full_name='kdeploy.AppConstructor.PortsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.AppConstructor.PortsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.AppConstructor.PortsEntry.value', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=428,
  serialized_end=472,
)

_APPCONSTRUCTOR = _descriptor.Descriptor(
  name='AppConstructor',
  full_name='kdeploy.AppConstructor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='kdeploy.AppConstructor.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='kdeploy.AppConstructor.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image', full_name='kdeploy.AppConstructor.image', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='env', full_name='kdeploy.AppConstructor.env', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='kdeploy.AppConstructor.ports', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='kdeploy.AppConstructor.replicas', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_APPCONSTRUCTOR_ENVENTRY, _APPCONSTRUCTOR_PORTSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=475,
  serialized_end=799,
)


_APPUPDATE_ENVENTRY = _descriptor.Descriptor(
  name='EnvEntry',
  full_name='kdeploy.AppUpdate.EnvEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.AppUpdate.EnvEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.AppUpdate.EnvEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=384,
  serialized_end=426,
)

_APPUPDATE_PORTSENTRY = _descriptor.Descriptor(
  name='PortsEntry',
  full_name='kdeploy.AppUpdate.PortsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='kdeploy.AppUpdate.PortsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='kdeploy.AppUpdate.PortsEntry.value', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=428,
  serialized_end=472,
)

_APPUPDATE = _descriptor.Descriptor(
  name='AppUpdate',
  full_name='kdeploy.AppUpdate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='kdeploy.AppUpdate.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='kdeploy.AppUpdate.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image', full_name='kdeploy.AppUpdate.image', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='env', full_name='kdeploy.AppUpdate.env', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='kdeploy.AppUpdate.ports', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='kdeploy.AppUpdate.replicas', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_APPUPDATE_ENVENTRY, _APPUPDATE_PORTSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=802,
  serialized_end=1093,
)


_APPREF = _descriptor.Descriptor(
  name='AppRef',
  full_name='kdeploy.AppRef',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='kdeploy.AppRef.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='kdeploy.AppRef.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342\337\037\014\n\n^.{1,225}$'), file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1095,
  serialized_end=1172,
)


_REPLICA = _descriptor.Descriptor(
  name='Replica',
  full_name='kdeploy.Replica',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='phase', full_name='kdeploy.Replica.phase', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='condition', full_name='kdeploy.Replica.condition', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reason', full_name='kdeploy.Replica.reason', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1174,
  serialized_end=1233,
)


_STATUS = _descriptor.Descriptor(
  name='Status',
  full_name='kdeploy.Status',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='replicas', full_name='kdeploy.Status.replicas', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1235,
  serialized_end=1279,
)


_LOG = _descriptor.Descriptor(
  name='Log',
  full_name='kdeploy.Log',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='kdeploy.Log.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1281,
  serialized_end=1303,
)


_APPS = _descriptor.Descriptor(
  name='Apps',
  full_name='kdeploy.Apps',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='applications', full_name='kdeploy.Apps.applications', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1305,
  serialized_end=1347,
)


_NAMESPACE = _descriptor.Descriptor(
  name='Namespace',
  full_name='kdeploy.Namespace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespace', full_name='kdeploy.Namespace.namespace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1349,
  serialized_end=1379,
)


_NAMESPACES = _descriptor.Descriptor(
  name='Namespaces',
  full_name='kdeploy.Namespaces',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaces', full_name='kdeploy.Namespaces.namespaces', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1381,
  serialized_end=1413,
)

_APP_ENVENTRY.containing_type = _APP
_APP_PORTSENTRY.containing_type = _APP
_APP.fields_by_name['env'].message_type = _APP_ENVENTRY
_APP.fields_by_name['ports'].message_type = _APP_PORTSENTRY
_APP.fields_by_name['status'].message_type = _STATUS
_APPCONSTRUCTOR_ENVENTRY.containing_type = _APPCONSTRUCTOR
_APPCONSTRUCTOR_PORTSENTRY.containing_type = _APPCONSTRUCTOR
_APPCONSTRUCTOR.fields_by_name['env'].message_type = _APPCONSTRUCTOR_ENVENTRY
_APPCONSTRUCTOR.fields_by_name['ports'].message_type = _APPCONSTRUCTOR_PORTSENTRY
_APPUPDATE_ENVENTRY.containing_type = _APPUPDATE
_APPUPDATE_PORTSENTRY.containing_type = _APPUPDATE
_APPUPDATE.fields_by_name['env'].message_type = _APPUPDATE_ENVENTRY
_APPUPDATE.fields_by_name['ports'].message_type = _APPUPDATE_PORTSENTRY
_STATUS.fields_by_name['replicas'].message_type = _REPLICA
_APPS.fields_by_name['applications'].message_type = _APP
DESCRIPTOR.message_types_by_name['App'] = _APP
DESCRIPTOR.message_types_by_name['AppConstructor'] = _APPCONSTRUCTOR
DESCRIPTOR.message_types_by_name['AppUpdate'] = _APPUPDATE
DESCRIPTOR.message_types_by_name['AppRef'] = _APPREF
DESCRIPTOR.message_types_by_name['Replica'] = _REPLICA
DESCRIPTOR.message_types_by_name['Status'] = _STATUS
DESCRIPTOR.message_types_by_name['Log'] = _LOG
DESCRIPTOR.message_types_by_name['Apps'] = _APPS
DESCRIPTOR.message_types_by_name['Namespace'] = _NAMESPACE
DESCRIPTOR.message_types_by_name['Namespaces'] = _NAMESPACES
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

App = _reflection.GeneratedProtocolMessageType('App', (_message.Message,), dict(

  EnvEntry = _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), dict(
    DESCRIPTOR = _APP_ENVENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.App.EnvEntry)
    ))
  ,

  PortsEntry = _reflection.GeneratedProtocolMessageType('PortsEntry', (_message.Message,), dict(
    DESCRIPTOR = _APP_PORTSENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.App.PortsEntry)
    ))
  ,
  DESCRIPTOR = _APP,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.App)
  ))
_sym_db.RegisterMessage(App)
_sym_db.RegisterMessage(App.EnvEntry)
_sym_db.RegisterMessage(App.PortsEntry)

AppConstructor = _reflection.GeneratedProtocolMessageType('AppConstructor', (_message.Message,), dict(

  EnvEntry = _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), dict(
    DESCRIPTOR = _APPCONSTRUCTOR_ENVENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.AppConstructor.EnvEntry)
    ))
  ,

  PortsEntry = _reflection.GeneratedProtocolMessageType('PortsEntry', (_message.Message,), dict(
    DESCRIPTOR = _APPCONSTRUCTOR_PORTSENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.AppConstructor.PortsEntry)
    ))
  ,
  DESCRIPTOR = _APPCONSTRUCTOR,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.AppConstructor)
  ))
_sym_db.RegisterMessage(AppConstructor)
_sym_db.RegisterMessage(AppConstructor.EnvEntry)
_sym_db.RegisterMessage(AppConstructor.PortsEntry)

AppUpdate = _reflection.GeneratedProtocolMessageType('AppUpdate', (_message.Message,), dict(

  EnvEntry = _reflection.GeneratedProtocolMessageType('EnvEntry', (_message.Message,), dict(
    DESCRIPTOR = _APPUPDATE_ENVENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.AppUpdate.EnvEntry)
    ))
  ,

  PortsEntry = _reflection.GeneratedProtocolMessageType('PortsEntry', (_message.Message,), dict(
    DESCRIPTOR = _APPUPDATE_PORTSENTRY,
    __module__ = 'kdeploy_pb2'
    # @@protoc_insertion_point(class_scope:kdeploy.AppUpdate.PortsEntry)
    ))
  ,
  DESCRIPTOR = _APPUPDATE,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.AppUpdate)
  ))
_sym_db.RegisterMessage(AppUpdate)
_sym_db.RegisterMessage(AppUpdate.EnvEntry)
_sym_db.RegisterMessage(AppUpdate.PortsEntry)

AppRef = _reflection.GeneratedProtocolMessageType('AppRef', (_message.Message,), dict(
  DESCRIPTOR = _APPREF,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.AppRef)
  ))
_sym_db.RegisterMessage(AppRef)

Replica = _reflection.GeneratedProtocolMessageType('Replica', (_message.Message,), dict(
  DESCRIPTOR = _REPLICA,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Replica)
  ))
_sym_db.RegisterMessage(Replica)

Status = _reflection.GeneratedProtocolMessageType('Status', (_message.Message,), dict(
  DESCRIPTOR = _STATUS,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Status)
  ))
_sym_db.RegisterMessage(Status)

Log = _reflection.GeneratedProtocolMessageType('Log', (_message.Message,), dict(
  DESCRIPTOR = _LOG,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Log)
  ))
_sym_db.RegisterMessage(Log)

Apps = _reflection.GeneratedProtocolMessageType('Apps', (_message.Message,), dict(
  DESCRIPTOR = _APPS,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Apps)
  ))
_sym_db.RegisterMessage(Apps)

Namespace = _reflection.GeneratedProtocolMessageType('Namespace', (_message.Message,), dict(
  DESCRIPTOR = _NAMESPACE,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Namespace)
  ))
_sym_db.RegisterMessage(Namespace)

Namespaces = _reflection.GeneratedProtocolMessageType('Namespaces', (_message.Message,), dict(
  DESCRIPTOR = _NAMESPACES,
  __module__ = 'kdeploy_pb2'
  # @@protoc_insertion_point(class_scope:kdeploy.Namespaces)
  ))
_sym_db.RegisterMessage(Namespaces)


DESCRIPTOR._options = None
_APP_ENVENTRY._options = None
_APP_PORTSENTRY._options = None
_APPCONSTRUCTOR_ENVENTRY._options = None
_APPCONSTRUCTOR_PORTSENTRY._options = None
_APPCONSTRUCTOR.fields_by_name['name']._options = None
_APPCONSTRUCTOR.fields_by_name['namespace']._options = None
_APPCONSTRUCTOR.fields_by_name['image']._options = None
_APPUPDATE_ENVENTRY._options = None
_APPUPDATE_PORTSENTRY._options = None
_APPUPDATE.fields_by_name['name']._options = None
_APPUPDATE.fields_by_name['namespace']._options = None
_APPREF.fields_by_name['name']._options = None
_APPREF.fields_by_name['namespace']._options = None

_KDEPLOYSERVICE = _descriptor.ServiceDescriptor(
  name='KdeployService',
  full_name='kdeploy.KdeployService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1416,
  serialized_end=1791,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateApp',
    full_name='kdeploy.KdeployService.CreateApp',
    index=0,
    containing_service=None,
    input_type=_APPCONSTRUCTOR,
    output_type=_APP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateApp',
    full_name='kdeploy.KdeployService.UpdateApp',
    index=1,
    containing_service=None,
    input_type=_APPUPDATE,
    output_type=_APP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteApp',
    full_name='kdeploy.KdeployService.DeleteApp',
    index=2,
    containing_service=None,
    input_type=_APPREF,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetApp',
    full_name='kdeploy.KdeployService.GetApp',
    index=3,
    containing_service=None,
    input_type=_APPREF,
    output_type=_APP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Logs',
    full_name='kdeploy.KdeployService.Logs',
    index=4,
    containing_service=None,
    input_type=_APPREF,
    output_type=_LOG,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListNamespaces',
    full_name='kdeploy.KdeployService.ListNamespaces',
    index=5,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_NAMESPACES,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListApps',
    full_name='kdeploy.KdeployService.ListApps',
    index=6,
    containing_service=None,
    input_type=_NAMESPACE,
    output_type=_APPS,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_KDEPLOYSERVICE)

DESCRIPTOR.services_by_name['KdeployService'] = _KDEPLOYSERVICE

# @@protoc_insertion_point(module_scope)

# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: unified.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\runified.proto\x12\nfileserver\"3\n\x0f\x46ileDescription\x12\x10\n\x08\x46ileName\x18\x01 \x01(\t\x12\x0e\n\x06\x42ucket\x18\x02 \x01(\t\":\n\x0cUploadStatus\x12*\n\x04\x43ode\x18\x01 \x01(\x0e\x32\x1c.fileserver.UploadStatusCode\"N\n\x05\x43hunk\x12\x34\n\x0f\x46ileDescription\x18\x01 \x01(\x0b\x32\x1b.fileserver.FileDescription\x12\x0f\n\x07\x43ontent\x18\x02 \x01(\x0c\"5\n\x10StructureRequest\x12\x0e\n\x06\x62ucket\x18\x01 \x01(\t\x12\x11\n\tdirectory\x18\x02 \x01(\t\"+\n\tDirectory\x12\x0c\n\x04Name\x18\x01 \x01(\t\x12\x10\n\x08\x46ileName\x18\x02 \x03(\t\"/\n\tStructure\x12\"\n\x06Object\x18\x01 \x03(\x0b\x32\x12.fileserver.Object\"$\n\x06Object\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\"G\n\x0f\x44ownloadRequest\x12\x34\n\x0f\x46ileDescription\x18\x01 \x01(\x0b\x32\x1b.fileserver.FileDescription\"\x07\n\x05\x45mpty\"\x14\n\x04Keys\x12\x0c\n\x04keys\x18\x01 \x01(\t\"&\n\x04User\x12\x0c\n\x04name\x18\x01 \x01(\x0c\x12\x10\n\x08password\x18\x02 \x01(\x0c\">\n\x05Token\x12\r\n\x05token\x18\x01 \x01(\t\x12&\n\x06status\x18\x02 \x01(\x0e\x32\x16.fileserver.AuthStatus\",\n\x0b\x41uthRequest\x12\r\n\x05token\x18\x01 \x01(\t\x12\x0e\n\x06\x61\x63\x63\x65ss\x18\x02 \x03(\t\"!\n\tAuthReply\x12\x14\n\x0cisAuthorized\x18\x01 \x01(\x08\"\x19\n\x05\x41\x64\x64\x65\x64\x12\x10\n\x08wasAdded\x18\x01 \x01(\x08*3\n\x10UploadStatusCode\x12\x0b\n\x07Unknown\x10\x00\x12\x06\n\x02Ok\x10\x01\x12\n\n\x06\x46\x61iled\x10\x02* \n\nAuthStatus\x12\x06\n\x02OK\x10\x00\x12\n\n\x06\x46\x41ILED\x10\x01\x32\xd0\x01\n\tRetriever\x12<\n\tsaveFiles\x12\x11.fileserver.Chunk\x1a\x18.fileserver.UploadStatus\"\x00(\x01\x12\x45\n\x0cgetStructure\x12\x1c.fileserver.StructureRequest\x1a\x15.fileserver.Structure\"\x00\x12>\n\x08getFiles\x12\x1b.fileserver.DownloadRequest\x1a\x11.fileserver.Chunk\"\x00\x30\x01\x32q\n\rauthenticator\x12\x30\n\x07getKeys\x12\x11.fileserver.Empty\x1a\x10.fileserver.Keys\"\x00\x12.\n\x05login\x12\x10.fileserver.User\x1a\x11.fileserver.Token\"\x00\x32\x90\x01\n\nauthorizer\x12@\n\x0cisAuthorized\x12\x17.fileserver.AuthRequest\x1a\x15.fileserver.AuthReply\"\x00\x12@\n\x10\x61\x64\x64\x41uthorization\x12\x17.fileserver.AuthRequest\x1a\x11.fileserver.Added\"\x00\x42\x0fZ\r./;fileserverb\x06proto3')

_UPLOADSTATUSCODE = DESCRIPTOR.enum_types_by_name['UploadStatusCode']
UploadStatusCode = enum_type_wrapper.EnumTypeWrapper(_UPLOADSTATUSCODE)
_AUTHSTATUS = DESCRIPTOR.enum_types_by_name['AuthStatus']
AuthStatus = enum_type_wrapper.EnumTypeWrapper(_AUTHSTATUS)
Unknown = 0
Ok = 1
Failed = 2
OK = 0
FAILED = 1


_FILEDESCRIPTION = DESCRIPTOR.message_types_by_name['FileDescription']
_UPLOADSTATUS = DESCRIPTOR.message_types_by_name['UploadStatus']
_CHUNK = DESCRIPTOR.message_types_by_name['Chunk']
_STRUCTUREREQUEST = DESCRIPTOR.message_types_by_name['StructureRequest']
_DIRECTORY = DESCRIPTOR.message_types_by_name['Directory']
_STRUCTURE = DESCRIPTOR.message_types_by_name['Structure']
_OBJECT = DESCRIPTOR.message_types_by_name['Object']
_DOWNLOADREQUEST = DESCRIPTOR.message_types_by_name['DownloadRequest']
_EMPTY = DESCRIPTOR.message_types_by_name['Empty']
_KEYS = DESCRIPTOR.message_types_by_name['Keys']
_USER = DESCRIPTOR.message_types_by_name['User']
_TOKEN = DESCRIPTOR.message_types_by_name['Token']
_AUTHREQUEST = DESCRIPTOR.message_types_by_name['AuthRequest']
_AUTHREPLY = DESCRIPTOR.message_types_by_name['AuthReply']
_ADDED = DESCRIPTOR.message_types_by_name['Added']
FileDescription = _reflection.GeneratedProtocolMessageType('FileDescription', (_message.Message,), {
  'DESCRIPTOR' : _FILEDESCRIPTION,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.FileDescription)
  })
_sym_db.RegisterMessage(FileDescription)

UploadStatus = _reflection.GeneratedProtocolMessageType('UploadStatus', (_message.Message,), {
  'DESCRIPTOR' : _UPLOADSTATUS,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.UploadStatus)
  })
_sym_db.RegisterMessage(UploadStatus)

Chunk = _reflection.GeneratedProtocolMessageType('Chunk', (_message.Message,), {
  'DESCRIPTOR' : _CHUNK,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Chunk)
  })
_sym_db.RegisterMessage(Chunk)

StructureRequest = _reflection.GeneratedProtocolMessageType('StructureRequest', (_message.Message,), {
  'DESCRIPTOR' : _STRUCTUREREQUEST,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.StructureRequest)
  })
_sym_db.RegisterMessage(StructureRequest)

Directory = _reflection.GeneratedProtocolMessageType('Directory', (_message.Message,), {
  'DESCRIPTOR' : _DIRECTORY,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Directory)
  })
_sym_db.RegisterMessage(Directory)

Structure = _reflection.GeneratedProtocolMessageType('Structure', (_message.Message,), {
  'DESCRIPTOR' : _STRUCTURE,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Structure)
  })
_sym_db.RegisterMessage(Structure)

Object = _reflection.GeneratedProtocolMessageType('Object', (_message.Message,), {
  'DESCRIPTOR' : _OBJECT,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Object)
  })
_sym_db.RegisterMessage(Object)

DownloadRequest = _reflection.GeneratedProtocolMessageType('DownloadRequest', (_message.Message,), {
  'DESCRIPTOR' : _DOWNLOADREQUEST,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.DownloadRequest)
  })
_sym_db.RegisterMessage(DownloadRequest)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), {
  'DESCRIPTOR' : _EMPTY,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Empty)
  })
_sym_db.RegisterMessage(Empty)

Keys = _reflection.GeneratedProtocolMessageType('Keys', (_message.Message,), {
  'DESCRIPTOR' : _KEYS,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Keys)
  })
_sym_db.RegisterMessage(Keys)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), {
  'DESCRIPTOR' : _USER,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.User)
  })
_sym_db.RegisterMessage(User)

Token = _reflection.GeneratedProtocolMessageType('Token', (_message.Message,), {
  'DESCRIPTOR' : _TOKEN,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Token)
  })
_sym_db.RegisterMessage(Token)

AuthRequest = _reflection.GeneratedProtocolMessageType('AuthRequest', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREQUEST,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.AuthRequest)
  })
_sym_db.RegisterMessage(AuthRequest)

AuthReply = _reflection.GeneratedProtocolMessageType('AuthReply', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREPLY,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.AuthReply)
  })
_sym_db.RegisterMessage(AuthReply)

Added = _reflection.GeneratedProtocolMessageType('Added', (_message.Message,), {
  'DESCRIPTOR' : _ADDED,
  '__module__' : 'unified_pb2'
  # @@protoc_insertion_point(class_scope:fileserver.Added)
  })
_sym_db.RegisterMessage(Added)

_RETRIEVER = DESCRIPTOR.services_by_name['Retriever']
_AUTHENTICATOR = DESCRIPTOR.services_by_name['authenticator']
_AUTHORIZER = DESCRIPTOR.services_by_name['authorizer']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\r./;fileserver'
  _UPLOADSTATUSCODE._serialized_start=725
  _UPLOADSTATUSCODE._serialized_end=776
  _AUTHSTATUS._serialized_start=778
  _AUTHSTATUS._serialized_end=810
  _FILEDESCRIPTION._serialized_start=29
  _FILEDESCRIPTION._serialized_end=80
  _UPLOADSTATUS._serialized_start=82
  _UPLOADSTATUS._serialized_end=140
  _CHUNK._serialized_start=142
  _CHUNK._serialized_end=220
  _STRUCTUREREQUEST._serialized_start=222
  _STRUCTUREREQUEST._serialized_end=275
  _DIRECTORY._serialized_start=277
  _DIRECTORY._serialized_end=320
  _STRUCTURE._serialized_start=322
  _STRUCTURE._serialized_end=369
  _OBJECT._serialized_start=371
  _OBJECT._serialized_end=407
  _DOWNLOADREQUEST._serialized_start=409
  _DOWNLOADREQUEST._serialized_end=480
  _EMPTY._serialized_start=482
  _EMPTY._serialized_end=489
  _KEYS._serialized_start=491
  _KEYS._serialized_end=511
  _USER._serialized_start=513
  _USER._serialized_end=551
  _TOKEN._serialized_start=553
  _TOKEN._serialized_end=615
  _AUTHREQUEST._serialized_start=617
  _AUTHREQUEST._serialized_end=661
  _AUTHREPLY._serialized_start=663
  _AUTHREPLY._serialized_end=696
  _ADDED._serialized_start=698
  _ADDED._serialized_end=723
  _RETRIEVER._serialized_start=813
  _RETRIEVER._serialized_end=1021
  _AUTHENTICATOR._serialized_start=1023
  _AUTHENTICATOR._serialized_end=1136
  _AUTHORIZER._serialized_start=1139
  _AUTHORIZER._serialized_end=1283
# @@protoc_insertion_point(module_scope)
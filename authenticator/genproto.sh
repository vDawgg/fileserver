#!/bin/bash

# 'import unified_pb2 as unified__pb2' needs to be changed to 'from . import unified_pb2 as unified__pb2' _pb2_grpc.py after gernerating!
python3 -m grpc_tools.protoc -I../proto --python_out=./proto --grpc_python_out=./proto ../proto/unified.proto

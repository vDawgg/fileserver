#!/bin/bash

# Proto files still need to be edited to work properly!
python -m grpc_tools.protoc -I../proto --python_out=proto --grpc_python_out=proto ../proto/unified.proto

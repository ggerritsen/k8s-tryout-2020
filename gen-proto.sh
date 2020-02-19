#!/bin/sh

protoc -I proto/ proto/greetsvc.proto --go_out=plugins=grpc:proto

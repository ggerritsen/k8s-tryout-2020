#!/bin/sh

set -e
set -o errexit

protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

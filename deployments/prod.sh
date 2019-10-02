#!/bin/bash -e

ARGS="--admin-bind-address=:8081"
# BEGIN __INCLUDE_GRPC__
ARGS="$ARGS --grpc-bind-address=:8080 --tracing-probability=0.01 --jaeger-host=localhost"
# END __INCLUDE_GRPC__

./go-template serve $ARGS

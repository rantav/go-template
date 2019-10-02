#!/bin/bash -e
AF_METRICS_TARGET_GRAPHITE_HOST=localhost

# BEGIN __INCLUDE_GRPC__
# Run Jaeger locally
(docker ps -f name=jaeger | grep jaeger) ||\
    ((docker rm jaeger || echo fine) &&\
        docker run -d --name jaeger \
          -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
          -p 5775:5775/udp \
          -p 6831:6831/udp \
          -p 6832:6832/udp \
          -p 5778:5778 \
          -p 16686:16686 \
          -p 14268:14268 \
          -p 9411:9411 \
          jaegertracing/all-in-one:1.8)
# END __INCLUDE_GRPC__
echo

echo
echo -e "\t \033[7mHealth:\033[0m http://localhost:8081/_/health.json"
echo -e "\t \033[7mMetrics:\033[0m http://localhost:8081/_/metrics.json"
# BEGIN __INCLUDE_GRPC__
echo -e "\t \033[7mrpcz:\033[0m http://localhost:8081/_/zpages/rpcz"
echo -e "\t \033[7mtracez:\033[0m http://localhost:8081/_/zpages/tracez"
echo -e "\t \033[7mchannelz:\033[0m http://localhost:8081/_/channelz/"
echo -e "\t \033[7mJaeger:\033[0m http://localhost:16686/"
# END __INCLUDE_GRPC__
echo

ARGS="--admin-bind-address=:8081"
# BEGIN __INCLUDE_GRPC__
ARGS="$ARGS --grpc-bind-address=:8080 --tracing-probability=1 --jaeger-host=localhost"
# END __INCLUDE_GRPC__

set -x
go run main.go serve $ARGS

package grpcclient

import (
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/encoding/gzip"

	grpclib "github.com/rantav/go-template/internal/generated/grpc"
)

// New creates a new grpc client for the API of GoTemplateClient
//
// connactionString could take several forms:
// * host:port
// * consul://127.0.0.1:8500/attr-endpoint?tag=attr-ep&timeout=2s&dc=us1
// * dns://127.0.0.1:8600/attr-endpoint.service.consul.:8080
//
// For a full spec of grpc connection strings (called NAMING) see
//		 https://github.com/grpc/grpc/blob/master/doc/naming.md
// For consul based naming see https://github.com/mbobakov/grpc-consul-resolver
func New(connectionString string) (grpclib.GoTemplateClient, error) {
	conn, err := grpc.Dial(connectionString,
		grpc.WithInsecure(), // Unfortunately for now...
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		grpc.WithStatsHandler(new(ocgrpc.ClientHandler)),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "error doaling to %s", connectionString)
	}

	client := grpclib.NewGoTemplateClient(conn)
	return client, err
}

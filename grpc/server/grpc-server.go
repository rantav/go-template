package grpcserver

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // make sure it's impored for side effect

	grpclib "github.com/rantav/go-template/internal/generated/grpc"
)

type server struct {
	baseLog zerolog.Context
}

// New creates a new grpc server
func New() (*grpc.Server, error) {
	lg := log.With().Str("context", "go-template")

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(new(ocgrpc.ServerHandler)),
	)

	server := &server{lg}

	grpclib.RegisterGoTemplateServer(grpcServer, server)
	return grpcServer, nil
}

func (s *server) Healthcheck(ctx context.Context, _ *empty.Empty) (
	*grpclib.HealthCheckResponse, error,
) {
	log.Debug().Msg("Healthcheck invoked. We're OK")
	return &grpclib.HealthCheckResponse{
		Status: grpclib.HealthCheckResponse_Serving,
	}, nil
}

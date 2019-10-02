package admin

import (
	"context"
	"time"

	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	grpclib "github.com/rantav/go-template/internal/generated/grpc"
)

// NewGrpcHealthcheck creates a new Check that when executed uses the client to determine health
func NewGrpcHealthcheck(client grpclib.GoTemplateClient, name string) checks.Check {
	lg := log.With().
		Str("context", "healthcheck").
		Str("healthcheckName", name).Logger()
	return grpcHealthcheck{
		name:   name,
		client: client,
		log:    lg,
	}
}

type grpcHealthcheck struct {
	name   string
	client grpclib.GoTemplateClient
	log    zerolog.Logger
}

// Execute the gRPC healthcheck
func (c grpcHealthcheck) Execute() (details interface{}, err error) {
	c.log.Debug().Msg("Running check...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.client.Healthcheck(ctx, &empty.Empty{})
}

// Name returns the name of the check
func (c grpcHealthcheck) Name() string {
	return c.name
}

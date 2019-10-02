package admin

import (
	"net"
	"time"

	health "github.com/AppsFlyer/go-sundheit"
	"github.com/rs/zerolog/log"

	grpcclient "github.com/rantav/go-template/grpc/client"
)

// CreateHealthchecks creates the healthckecks and schedules them to run
func CreateHealthchecks(
	// BEGIN __INCLUDE_GRPC__
	grpcListenerAddress net.Addr,
	// END __INCLUDE_GRPC__
) *health.Health {
	lg := log.With().Str("context", "healthcheck").Logger()
	lg.Info().Msg("starting healthcheck server")

	// create a new health instance
	h := health.New()

	// BEGIN __INCLUDE_GRPC__
	// Schedule a gRPC healthcheck to our local gRPC service
	localGrpcClient, err := grpcclient.New(grpcListenerAddress.String())
	if err != nil {
		lg.Fatal().Msgf("Error connecting to local gRPC service. %+v", err)
	}
	err = h.RegisterCheck(&health.Config{
		Check:           NewGrpcHealthcheck(localGrpcClient, "localhost-grpc"),
		ExecutionPeriod: 1 * time.Minute,
	})
	if err != nil {
		lg.Fatal().Msgf("Error registering grpc healthcheck. %+v", err)
	}
	// END __INCLUDE_GRPC__

	// TODO: Various healthchecks will go here...
	// Remote gRPC servers (use NewGrpcHealthcheck)
	// Databases
	// Kafkas
	// Local files
	// etc

	return &h
}

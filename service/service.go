package service

import (
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	channelz "github.com/rantav/go-grpc-channelz"
	channelzservice "google.golang.org/grpc/channelz/service"

	"github.com/rantav/go-template/admin"
	grpcserver "github.com/rantav/go-template/grpc/server"
	"github.com/rantav/go-template/types"
)

// Serve starts the server serving
func Serve(
	adminBindAddress types.BindAddress,
	// BEGIN __INCLUDE_GRPC__
	grpcBindAddress types.BindAddress,
	tracingProbability float64,
	jaegerHost string,
	// END __INCLUDE_GRPC__
) {
	// BEGIN __INCLUDE_GRPC__
	grpcListener, err := net.Listen("tcp", string(grpcBindAddress))
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	grpcListenerAddress := grpcListener.Addr()

	_, err = admin.ConfigureTracing(tracingProbability, jaegerHost)
	if err != nil {
		log.Fatal().Msgf("Failed to configure tracing %+v", err)
	}

	// grpc server setup
	grpcServer, err := grpcserver.New()
	if err != nil {
		log.Fatal().Msgf("Failed to create grpc server %+v", err)
	}

	// Register the channelz handler
	http.Handle("/", channelz.CreateHandler("/_", string(grpcBindAddress)))
	// Register the channelz gRPC service to grpcServer so that we can query it for this service.
	channelzservice.RegisterChannelzServiceToServer(grpcServer)
	// END __INCLUDE_GRPC__

	// admin
	adminListener, err := net.Listen("tcp", string(adminBindAddress))
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}
	healthckecks := admin.CreateHealthchecks(
		// BEGIN __INCLUDE_GRPC__
		grpcListenerAddress,
		// END __INCLUDE_GRPC__
	)
	admin.RegisterHandlers(healthckecks)

	g := new(errgroup.Group)
	// BEGIN __INCLUDE_GRPC__
	g.Go(func() error { return grpcServer.Serve(grpcListener) })
	// END __INCLUDE_GRPC__
	g.Go(func() error { return http.Serve(adminListener, nil) })

	log.Info().Msgf("go-template is up; admin bind address: %s", adminBindAddress)
	log.Fatal().Msgf("Error running server: %s", g.Wait())
}

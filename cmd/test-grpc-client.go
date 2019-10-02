package cmd

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/cobra"

	"github.com/rs/zerolog/log"

	"github.com/rantav/go-template/admin"
	grpcclient "github.com/rantav/go-template/grpc/client"
)

var (
	testGrpcClientServerAddress *string
)

// serveCmd represents the serve command
var testGrpcClientCmd = &cobra.Command{
	Use:   "test-grpc-client",
	Short: "Test the gRPC client, connecting to a running server",
	Long: `First run the server using the "serve" command.
Then run this test-grpc-client command to connect to this server and test it.`,
	Run: func(cmd *cobra.Command, args []string) {
		exporter, err := admin.ConfigureTracing(1, "localhost")
		if err != nil {
			log.Error().Msgf("Error configuring tracing %+v", err)
		}
		defer exporter.Flush()

		client, err := grpcclient.New(*testGrpcClientServerAddress)
		if err != nil {
			log.Fatal().Msgf("Cannot create gRPC client to %s. %v", *testGrpcClientServerAddress, err)
		}
		resp, err := client.Healthcheck(context.Background(), &empty.Empty{})
		if err != nil {
			log.Error().Msgf("Error running healthcheck. %+v", err)
			return
		}
		log.Info().Msgf("Healthcheck result: %s", resp)
	},
}

func init() {
	rootCmd.AddCommand(testGrpcClientCmd)
	testGrpcClientServerAddress = addRequiredStringFlag(testGrpcClientCmd, "server-address", "",
		"Network address of the server. Could use simple host:port as well as DNS or consul address")
}

package cmd

import (
	"github.com/spf13/cobra"

	// BEGIN __INCLUDE_GRPC__
	"github.com/rantav/go-template/admin"
	// END __INCLUDE_GRPC__
	"github.com/rantav/go-template/service"
	"github.com/rantav/go-template/types"
)

var (
	adminBindAddress *string
	// BEGIN __INCLUDE_GRPC__
	grpcBindAddress    *string
	tracingProbability *float64
	jaegerHost         *string
	// END __INCLUDE_GRPC__
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run this service",
	Run: func(cmd *cobra.Command, args []string) {
		service.Serve(
			types.BindAddress(*adminBindAddress),
			// BEGIN __INCLUDE_GRPC__
			types.BindAddress(*grpcBindAddress),
			*tracingProbability,
			*jaegerHost,
			// END __INCLUDE_GRPC__
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	adminBindAddress = addRequiredStringFlag(serveCmd, "admin-bind-address", "",
		"Network address to bind to. For ex :8081 or 127.0.0.1:8081")
	// BEGIN __INCLUDE_GRPC__
	grpcBindAddress = addRequiredStringFlag(serveCmd, "grpc-bind-address", "",
		"Network address to bind to. For ex :8080 or 127.0.0.1:8080")
	tracingProbability = serveCmd.Flags().Float64("tracing-probability",
		admin.TracingDefaultProbability,
		"Tracing probability for each request")
	jaegerHost = addRequiredStringFlag(serveCmd, "jaeger-host", "",
		"Hosotname of the the jaeger tracing agent")
	// END __INCLUDE_GRPC__
}

package admin

import (
	"fmt"

	"contrib.go.opencensus.io/exporter/jaeger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"

	"github.com/rantav/go-template/version"
)

const TracingDefaultProbability = 0.005

const (
	// Port details: https://www.jaegertracing.io/docs/getting-started/
	jaegerAgentCompactThriftPort = 6831
	jaegerCollectorThriftPort    = 14268
)

// ConfigureTracing configures tracing.
// If jaegerHost is not provided, traces will remain in memory and not sent
func ConfigureTracing(tracingProbability float64, jaegerHost string) (*jaeger.Exporter, error) {
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(tracingProbability)})
	if jaegerHost == "" {
		log.Info().Msg("Jaeger host not defiend. Not registering a Jaeger exporter.")
		return nil, nil
	}
	return registerJaegerExporter(jaegerHost)
}

func registerJaegerExporter(jaegerHost string) (*jaeger.Exporter, error) {
	agentEndpointURI := fmt.Sprintf("%s:%d", jaegerHost, jaegerAgentCompactThriftPort)
	collectorEndpointURI := fmt.Sprintf("http://%s:%d/api/traces", jaegerHost, jaegerCollectorThriftPort)

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:     agentEndpointURI,
		CollectorEndpoint: collectorEndpointURI,
		ServiceName:       version.ServiceName,
		Process: jaeger.Process{
			ServiceName: version.ServiceName,
			Tags: []jaeger.Tag{
				jaeger.StringTag("revision", version.GitHash),
				jaeger.StringTag("branch", version.GitBranch),
				jaeger.StringTag("tag", version.GitTag),
				jaeger.StringTag("commit-message", version.GitCommitMessage),
				jaeger.StringTag("build-time", version.BuildTime),
			},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating Jaeger exporter")
	}

	// And now finally register it as a Trace Exporter
	trace.RegisterExporter(je)

	log.Info().Msgf("Jaeger host configured for tracing: %s", jaegerHost)
	return je, nil
}

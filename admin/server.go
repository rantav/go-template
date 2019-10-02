package admin

import (
	"net/http"

	// BEGIN __INCLUDE_GRPC__
	"go.opencensus.io/zpages"
	// END __INCLUDE_GRPC__

	health "github.com/AppsFlyer/go-sundheit"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
)

// RegisterHandlers registers the admin HTTP handlers
//
// Visit healthchecks at http://localhost:port/_/health.json
// View metrics at /_/metrics.json
// Unfortunately at the moment the metrics from opencensus are not yet visible at metrics.json
func RegisterHandlers(healthchecks *health.Health) {
	http.Handle("/_/health.json", healthhttp.HandleHealthJSON(*healthchecks))
	// BEGIN __INCLUDE_GRPC__

	// zpages open a UI for gRPC tracing.
	// See http://localhost:8080/_/zpages/rpcz
	// And http://localhost:8080/_/zpages/tracez
	zpages.Handle(http.DefaultServeMux, "/_/zpages")
	// END __INCLUDE_GRPC__
}

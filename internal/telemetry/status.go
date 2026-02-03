package telemetry

import (
	"net/http"
)

// HealthHandler is a simple HTTP handler that returns a 200 OK status.
// It can be used for health checks.
func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

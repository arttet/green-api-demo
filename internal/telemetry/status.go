package telemetry

import (
	"net/http"
)

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}

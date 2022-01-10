package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsEndpoint() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)
}

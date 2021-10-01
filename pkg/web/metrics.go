package web

import "github.com/gorilla/mux"
import "github.com/prometheus/client_golang/prometheus/promhttp"

func registerMetricHandler(r *mux.Router) {
	r.Path("/metrics").Handler(promhttp.Handler())
}

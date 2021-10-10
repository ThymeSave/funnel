package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/thymesave/funnel/pkg/health"
	"net/http"
)

type healthResponse struct {
	Status health.Status `json:"status"`
}

// HealthHandler reports the health status for the entire application or a specific component
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	component, componentBased := vars["component"]

	var status health.Status

	if componentBased {
		check := health.GetCheck(component)
		if check == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		checkFunc := *check
		status = checkFunc(context.Background())
	} else {
		status = health.GetAggregatedStatus(context.Background())
	}

	var httpStatus int
	if status == health.StatusHealthy {
		httpStatus = http.StatusOK
	} else {
		httpStatus = http.StatusInternalServerError
	}

	_ = SendJSON(w, httpStatus, healthResponse{Status: status})
}

package health

import (
	"context"
	"github.com/thymesave/funnel/pkg/config"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func mockResponse(httpStatus int, responseBody string, testFunc func(*httptest.Server)) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(httpStatus)
		_, _ = w.Write([]byte(responseBody))
	}))
	testFunc(server)
	defer server.Close()
}

func TestRegisterCheck(t *testing.T) {
	RegisterCheck("couchdb", func(ctx context.Context) Status {
		return StatusHealthy
	})
	if checks["couchdb"] == nil {
		t.Fatal("Expected health-check to be added to available checks")
	}
}

func TestGetCheck(t *testing.T) {
	if c := GetCheck("notFound"); c != nil {
		t.Fatalf("Expected check to be nil if none is registered")
	}

	RegisterCheck("foo", func(ctx context.Context) Status {
		return StatusHealthy
	})

	if c := GetCheck("foo"); c == nil {
		t.Fatalf("Expected check to be not nil")
	}
}

func TestGetAggregatedStatus(t *testing.T) {
	testCases := []struct {
		status           []Status
		aggregatedStatus Status
	}{
		{
			status:           []Status{StatusHealthy, StatusHealthy},
			aggregatedStatus: StatusHealthy,
		},
		{
			status:           []Status{StatusHealthy, StatusHealthy, StatusUnhealthy},
			aggregatedStatus: StatusUnhealthy,
		},
		{
			status:           []Status{},
			aggregatedStatus: StatusHealthy,
		},
		{
			status:           []Status{StatusUnhealthy},
			aggregatedStatus: StatusUnhealthy,
		},
	}

	for _, tc := range testCases {
		checks = map[string]CheckFunc{}
		// register checks
		for i, status := range tc.status {
			RegisterCheck(strconv.Itoa(i), func(ctx context.Context) Status {
				return status
			})
		}

		if status := GetAggregatedStatus(context.Background()); status != tc.aggregatedStatus {
			t.Fatalf("Expected aggregated status %s for inputs %v, but got %s", tc.aggregatedStatus, tc.status, status)
		}
	}
}

func init() {
	config.CreateDefault()
}

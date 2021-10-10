package health

import (
	"context"
	"github.com/thymesave/funnel/pkg/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCouchDBCheck(t *testing.T) {
	testCases := []struct {
		upstreamStatus int
		expectedStatus Status
		closeServer    bool
	}{
		{
			upstreamStatus: http.StatusOK,
			expectedStatus: StatusHealthy,
		},
		{
			upstreamStatus: http.StatusNotFound,
			expectedStatus: StatusUnhealthy,
		},
		{
			upstreamStatus: http.StatusInternalServerError,
			expectedStatus: StatusUnhealthy,
		},
		{
			upstreamStatus: -1,
			expectedStatus: StatusUnhealthy,
			closeServer:    true,
		},
	}

	for _, tc := range testCases {
		mockResponse(tc.upstreamStatus, "", func(server *httptest.Server) {
			if tc.closeServer {
				server.Close()
			}

			config.Get().CouchDB.ParseEndpoint(server.URL)

			if status := CouchDBCheck(context.Background()); status != tc.expectedStatus {
				t.Fatalf("Expected %s, but got %s for upstream http code %d", tc.expectedStatus, status, tc.upstreamStatus)
			}
		})
	}
}

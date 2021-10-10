package web

import (
	"net/http"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	testCases := []struct {
		expectedStatus int
		path           string
	}{
		{
			expectedStatus: http.StatusNotFound,
			path:           "/health/non-existing-component",
		},
		{
			expectedStatus: http.StatusOK,
			path:           "/health/general",
		},
		{
			expectedStatus: http.StatusInternalServerError,
			path:           "/health",
		},
	}

	for _, tc := range testCases {
		rr := testHandler(createRouter(), "GET", tc.path, nil)
		if rr.Code != tc.expectedStatus {
			t.Fatalf("Expected http status %d, but got %d for path %s", tc.expectedStatus, rr.Code, tc.path)
		}
	}
}

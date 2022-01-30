package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCORSProxyHandler(t *testing.T) {
	os.Setenv("FUNNEL_CORS_FORCE_LOCAL_RESOLUTION", "true")

	testCases := []struct {
		upstreamStatus int
		expectedStatus int
		contentType    string
		closeServer    bool
		useValidURL    bool
	}{
		{
			upstreamStatus: 200,
			expectedStatus: 200,
			contentType:    "text/html",
			useValidURL:    true,
		},
		{
			upstreamStatus: 200,
			expectedStatus: 200,
			contentType:    "text/html; charset=utf8",
			useValidURL:    true,
		},
		{
			upstreamStatus: 200,
			expectedStatus: 400,
			contentType:    "application/json",
			useValidURL:    true,
		},
		{
			upstreamStatus: -1,
			expectedStatus: 502,
			closeServer:    true,
			useValidURL:    true,
		},
		{
			upstreamStatus: -1,
			expectedStatus: 400,
			closeServer:    false,
			useValidURL:    false,
		},
	}

	for _, tc := range testCases {
		mockResponseWithContentType(tc.upstreamStatus, "", tc.contentType, func(server *httptest.Server) {
			if tc.closeServer {
				server.Close()
			}
			url := ""
			if tc.useValidURL {
				url = server.URL
			}

			req, _ := http.NewRequest("GET", "/service/cors-proxy/?url="+url, nil)
			rr := testHandlerWithRequest(http.HandlerFunc(CORSProxyHandler), req)
			if rr.Code != tc.expectedStatus {
				t.Fatalf("Expected status code %d but got %d with content type %s", tc.expectedStatus, rr.Code, tc.contentType)
			}
		})
	}
	os.Setenv("FUNNEL_CORS_FORCE_LOCAL_RESOLUTION", "false")
}

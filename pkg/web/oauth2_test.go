package web

import (
	"context"
	"net/http"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
)

func TestCreateOAuth2Handler(t *testing.T) {
	err := patchOidConfig("oidc-config")
	if err != nil {
		t.Fatal(err)
	}

	handler, err := CreateOAuth2Handler(oidc.InsecureIssuerURLContext(context.Background(), "https://auth.provider"))
	if err != nil {
		t.Fatal("Expected dummy provider config to be used", err)
	}

	testCases := []struct {
		authHeader     string
		expectedStatus int
	}{
		{
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			authHeader:     "Bearer",
			expectedStatus: http.StatusBadRequest,
		},
		{
			authHeader:     "Bearer abc",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		req, _ := http.NewRequest("GET", "baz", nil)
		if tc.authHeader != "" {
			req.Header.Set("Authorization", tc.authHeader)
		}
		rr := testHandlerWithRequest(handler(func(writer http.ResponseWriter, request *http.Request) {}), req)

		if rr.Code != tc.expectedStatus {
			t.Fatalf("Expected status code %d, but got %d", tc.expectedStatus, rr.Code)
		}
	}
}

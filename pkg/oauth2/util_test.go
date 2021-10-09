package oauth2

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/thymesave/funnel/pkg/config"
)

func TestAddTokenToRequestContext(t *testing.T) {
	r := (&http.Request{}).WithContext(context.Background())
	ctx := AddTokenToRequestContext(r, &oidc.IDToken{Subject: "test"})
	if ctx == nil {
		t.Fatal("Expected context to be not nil after modfication")
	}
}

func TestGetTokenFromRequest(t *testing.T) {
	r := (&http.Request{}).WithContext(context.Background())
	ctx := AddTokenToRequestContext(r, &oidc.IDToken{Subject: "test"})
	r = r.WithContext(ctx)

	token := GetTokenFromRequest(r)
	if token == nil {
		t.Fatal("Expected token to be present in context after setting")
	}
}

func TestExtractClaims(t *testing.T) {
	testCases := []struct {
		rawClaims      map[string]string
		names          []string
		expectedClaims map[string]string
		errorExpected  bool
	}{
		{
			rawClaims:      map[string]string{"username": "test"},
			names:          []string{"username"},
			expectedClaims: map[string]string{"username": "test"},
			errorExpected:  false,
		},
		{
			rawClaims:      map[string]string{"username": "test"},
			names:          []string{"username", "nonExistent"},
			expectedClaims: map[string]string{},
			errorExpected:  true,
		},
		{
			rawClaims:      map[string]string{"username": "test", "some": "val"},
			names:          []string{"username"},
			expectedClaims: map[string]string{"username": "test"},
			errorExpected:  false,
		},
	}

	for _, tc := range testCases {
		claims, err := ExtractClaims(tc.rawClaims, tc.names)
		if err != nil && !tc.errorExpected {
			t.Fatalf("Expected not to fail, but failed with %s", err)
		} else if err == nil && tc.errorExpected {
			t.Fatalf("Expected to fail, but didn't")
		}

		if (len(claims) != 0 && len(tc.expectedClaims) != 0) && !reflect.DeepEqual(claims, tc.expectedClaims) {
			t.Fatalf("Expected claims to match, expected %v, got %v", tc.expectedClaims, claims)
		}
	}
}

func TestGetUsername(t *testing.T) {
	oauth2Config := config.Get().Oauth2

	r := (&http.Request{}).WithContext(context.Background())

	// Try without token
	if _, err := GetUsername(r, oauth2Config); err == nil {
		t.Fatal("Expected empty context to throw error")
	}

	// Try with token without claims
	ctx := AddTokenToRequestContext(r, &oidc.IDToken{Subject: "test"})
	r = r.WithContext(ctx)

	username, err := GetUsername(r, oauth2Config)
	if err == nil {
		t.Fatal("Expected non existent claim to fail")
	}
	if username != "" {
		t.Fatal("Expected username to be blank")
	}
}

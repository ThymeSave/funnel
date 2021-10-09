package oauth2

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/thymesave/funnel/pkg/config"
)

func TestAddTokenToRequestContext(t *testing.T) {
	r := (&http.Request{}).WithContext(context.Background())
	ctx := AddTokenToRequestContext(r, NewIDToken("https://auth.provider", "test"))
	if ctx == nil {
		t.Fatal("Expected context to be not nil after modfication")
	}
}

func TestGetTokenFromRequest(t *testing.T) {
	r := (&http.Request{}).WithContext(context.Background())
	ctx := AddTokenToRequestContext(r, NewIDToken("https://auth.provider", "test"))
	r = r.WithContext(ctx)

	token := GetTokenFromRequest(r)
	if token == nil {
		t.Fatal("Expected token to be present in context after setting")
	}
}

func TestExtractClaims(t *testing.T) {
	testCases := []struct {
		rawClaims      map[string]interface{}
		names          []string
		expectedClaims map[string]interface{}
		errorExpected  bool
	}{
		{
			rawClaims:      map[string]interface{}{"username": "test"},
			names:          []string{"username"},
			expectedClaims: map[string]interface{}{"username": "test"},
			errorExpected:  false,
		},
		{
			rawClaims:      map[string]interface{}{"username": "test"},
			names:          []string{"username", "nonExistent"},
			expectedClaims: map[string]interface{}{},
			errorExpected:  true,
		},
		{
			rawClaims:      map[string]interface{}{"username": "test", "some": "val"},
			names:          []string{"username"},
			expectedClaims: map[string]interface{}{"username": "test"},
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
	oauth2Config.UsernameClaim = "email"

	r := (&http.Request{}).WithContext(context.Background())

	// Try without token
	if _, err := GetUsername(r, oauth2Config); err == nil {
		t.Fatal("Expected empty context to throw error")
	}

	// Try with token without claims
	ctx := AddTokenToRequestContext(r, NewIDToken("https://auth.provider", "test"))
	r = r.WithContext(ctx)

	username, err := GetUsername(r, oauth2Config)
	if err == nil {
		t.Fatal("Expected non existent claim to fail")
	}
	if username != "" {
		t.Fatal("Expected username to be blank")
	}

	// Try with valid token
	idToken := NewIDToken("https://auth.provider", "test")
	idToken.Claims = map[string]interface{}{"email": "test"}
	ctx = AddTokenToRequestContext(r, idToken)
	r = r.WithContext(ctx)
	username, err = GetUsername(r, oauth2Config)
	if err != nil {
		t.Fatalf("Expected error to be nil when username is present in claims, but got %s", err)
	}
	if username != "test" {
		t.Fatalf("Expected username to be test, but got %s", username)
	}
}

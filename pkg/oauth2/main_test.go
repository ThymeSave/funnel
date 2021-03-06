package oauth2

import (
	"context"
	"os"
	"testing"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/thymesave/funnel/pkg/config"
)

func init() {
	os.Setenv("FUNNEL_OAUTH2_ISSUER_URL", "")
	os.Setenv("FUNNEL_OAUTH2_CLIENT_ID", "")
	config.CreateDefault()
}

func patchOidConfig(configPath string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// Set to current package and testdata folder
	config.Get().Oauth2.IssuerURL = "file://" + path + "/testdata/" + configPath
	return nil
}

func createProvider(configPath string, issuerURL string) (*oidc.Provider, error) {
	err := patchOidConfig(configPath)
	if err != nil {
		return nil, err
	}
	return NewProvider(context.Background(), config.Get().Oauth2)
}

func TestNewProvider(t *testing.T) {
	_, err := createProvider("valid-oauth2", "https://auth.provider")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewVerifier(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		configPath string
	}{
		{
			shouldFail: false,
			configPath: "valid-oauth2",
		},
		{
			shouldFail: true,
			configPath: "non-existent",
		},
	}

	for _, testCase := range testCases {
		err := patchOidConfig(testCase.configPath)
		if err != nil {
			t.Fatal(err)
		}
		verifier, err := NewVerifier(context.Background(), config.Get().Oauth2)
		if err != nil {
			return
		}
		if testCase.shouldFail && err == nil {
			t.Fatalf("Expected verifier creation to fail with configPath=%s", testCase.configPath)
		} else if !testCase.shouldFail && (err != nil || verifier == nil) {
			t.Fatalf("Expected verifier creation NOT to fail with configPath=%s, but it returned a valid provider", testCase.configPath)
		}
	}
}

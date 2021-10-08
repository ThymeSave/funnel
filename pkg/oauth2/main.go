package oauth2

import (
	"context"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/thymesave/funnel/pkg/config"
)

type contextKey string

// MaxConnections specifies the amount of maximum connections allowed
const MaxConnections = 1000

// RequestTimeoutSeconds specified the amount of seconds to wait before canceling a http request for oauth2
const RequestTimeoutSeconds = 3

// TokenContextKey is the context key where the raw oidc token is stored
const TokenContextKey = contextKey("oauth2Token")

func createHTTPTransport() *http.Transport {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	t.MaxIdleConns = MaxConnections
	t.MaxConnsPerHost = MaxConnections
	t.MaxIdleConnsPerHost = MaxConnections
	return t
}

// NewProvider generates a new provider based on given configuration and modifies the default http client before doing so
func NewProvider(ctx context.Context, oauth2Config *config.OAuth2) (*oidc.Provider, error) {
	httpClient := &http.Client{
		Timeout:   RequestTimeoutSeconds * time.Second,
		Transport: createHTTPTransport(),
	}

	if !oauth2Config.VerifyIssuer {
		ctx = oidc.InsecureIssuerURLContext(ctx, oauth2Config.IssuerURL)
	}

	return oidc.NewProvider(oidc.ClientContext(ctx, httpClient), oauth2Config.IssuerURL)
}

// NewVerifier creates a new oauth2 verifier based on the given configuration
func NewVerifier(ctx context.Context, oauth2Config *config.OAuth2) (*oidc.IDTokenVerifier, error) {
	provider, err := NewProvider(ctx, oauth2Config)
	if err != nil {
		return nil, err
	}

	cfg := oidc.Config{
		ClientID: oauth2Config.ClientID,
	}

	return provider.Verifier(&cfg), nil
}

// AddTokenToRequestContext adds the given token as context field
func AddTokenToRequestContext(r *http.Request, token *oidc.IDToken) context.Context {
	return context.WithValue(r.Context(), TokenContextKey, token)
}

// GetTokenFromRequest return ths given token from context
func GetTokenFromRequest(r *http.Request) *oidc.IDToken {
	return r.Context().Value(TokenContextKey).(*oidc.IDToken)
}

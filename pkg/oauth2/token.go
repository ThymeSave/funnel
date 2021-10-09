package oauth2

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"time"
)

// IDToken is a wrapper around oidc.IDToken to make it easier to test and use
type IDToken struct {
	raw    *oidc.IDToken
	Claims map[string]interface{}
}

// Convert the given oidc token
func Convert(token *oidc.IDToken) *IDToken {
	ourToken := IDToken{raw: token}
	// Ignore errors
	_ = token.Claims(&ourToken.Claims)
	return &ourToken
}

// NewIDToken generates a new id token and is intended to be used for testing
func NewIDToken(issuer string, clientID string) *IDToken {
	return Convert(&oidc.IDToken{
		Issuer:          issuer,
		Audience:        []string{clientID},
		Subject:         "sub",
		Expiry:          time.Now().Add(1 * time.Hour),
		IssuedAt:        time.Now(),
		Nonce:           "",
		AccessTokenHash: "",
	})
}

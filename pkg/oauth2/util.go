package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/thymesave/funnel/pkg/config"
)

// AddTokenToRequestContext adds the given token as context field
func AddTokenToRequestContext(r *http.Request, token *oidc.IDToken) context.Context {
	return context.WithValue(r.Context(), TokenContextKey, token)
}

// GetTokenFromRequest return ths given token from context
func GetTokenFromRequest(r *http.Request) *oidc.IDToken {
	val := r.Context().Value(TokenContextKey)
	if val == nil {
		return nil
	}
	return val.(*oidc.IDToken)
}

func getClaims(r *http.Request, names []string) (map[string]string, error) {
	token := GetTokenFromRequest(r)
	if token == nil {
		return nil, ErrNoTokenInContext
	}

	rawClaims := make(map[string]string, 0)

	err := token.Claims(&rawClaims)
	if err != nil {
		return nil, err
	}

	return ExtractClaims(rawClaims, names)
}

// ExtractClaims from the given claims
func ExtractClaims(rawClaims map[string]string, names []string) (map[string]string, error) {
	claims := make(map[string]string, len(names))
	for _, name := range names {
		val, ok := rawClaims[name]
		if !ok {
			return nil, fmt.Errorf("claim %s is not present", name)
		}
		claims[name] = val
	}

	return claims, nil
}

// GetUsername from request context, enhanced by oauth2 middleware
func GetUsername(r *http.Request, oauth2Config *config.OAuth2) (string, error) {
	usernameClaim := oauth2Config.UsernameClaim
	claims, err := getClaims(r, []string{usernameClaim})
	if err != nil {
		return "", err
	}

	return claims[usernameClaim], nil
}

package web

import (
	"context"
	"net/http"
	"strings"

	"github.com/thymesave/funnel/pkg/config"
	"github.com/thymesave/funnel/pkg/oauth2"
)

// AuthErrorResponse is the representation of an oauth error
type AuthErrorResponse struct {
	Message string `json:"message"`
}

func oauthError(w http.ResponseWriter, status int, response AuthErrorResponse) {
	w.WriteHeader(status)
	_ = SendJSON(w, AuthErrorResponse{Message: "Authorization header is missing"})
}

// CreateOAuth2Handler in given context, creates an oauth2 verifier, and returns a middleware
func CreateOAuth2Handler(ctx context.Context) (func(next http.HandlerFunc) http.HandlerFunc, error) {
	verifier, err := oauth2.NewVerifier(ctx, config.Get().Oauth2)
	if err != nil {
		return nil, err
	}

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Get auth header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				oauthError(w, http.StatusUnauthorized, AuthErrorResponse{Message: "Authorization header is missing"})
				return
			}

			// Try to extract bearer part
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 {
				oauthError(w, http.StatusBadRequest, AuthErrorResponse{Message: "Authorization header format is invalid"})
				return
			}

			// Verify JWT is signed and valid
			idToken, err := verifier.Verify(ctx, parts[1])
			if err != nil {
				oauthError(w, http.StatusUnauthorized, AuthErrorResponse{Message: "Failed to validate idToken: " + err.Error()})
				return
			}

			next(w, r.WithContext(oauth2.AddTokenToRequestContext(r, idToken)))
		}
	}, nil
}

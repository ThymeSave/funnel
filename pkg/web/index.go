package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/thymesave/funnel/pkg/config"

	"github.com/thymesave/funnel/pkg/buildinfo"
)

// IndexHandler is the default response a curious user will see
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	oauth2Config := config.Get().Oauth2

	_ = SendJSON(w, map[string]interface{}{
		"funnel": map[string]string{
			"version":   buildinfo.Version,
			"gitSha":    buildinfo.GitSha,
			"buildTime": buildinfo.BuildTime,
		},
		"oidc": map[string]interface{}{
			"issuerUrl":     oauth2Config.IssuerURL,
			"configUrl":     fmt.Sprintf("%s/.well-known/openid-configuration", strings.TrimSuffix(oauth2Config.IssuerURL, "/")),
			"clientId":      oauth2Config.ClientID,
			"scopes":        oauth2Config.Scopes,
			"usernameClaim": oauth2Config.UsernameClaim,
		},
	})
}

package web

import (
	"net/http"

	"github.com/thymesave/funnel/pkg/buildinfo"
)

// IndexHandler is the default response a curious user will see
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_ = SendJSON(w, map[string]interface{}{
		"info": map[string]string{
			"version":   buildinfo.Version,
			"gitSha":    buildinfo.GitSha,
			"buildTime": buildinfo.BuildTime,
		},
	})
}

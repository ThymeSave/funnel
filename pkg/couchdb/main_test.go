package couchdb

import (
	"github.com/thymesave/funnel/pkg/config"
	"os"
)

func init() {
	os.Setenv("FUNNEL_OAUTH2_ISSUER_URL", "")
	os.Setenv("FUNNEL_OAUTH2_CLIENT_ID", "")
	// Create default configs for tests
	config.CreateDefault()
}

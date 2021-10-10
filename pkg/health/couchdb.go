package health

import (
	"context"
	"fmt"
	"github.com/thymesave/funnel/pkg/config"
	"net/http"
)

// CouchDBCheck verifies couchdb is reachable
func CouchDBCheck(ctx context.Context) Status {
	endpoint := fmt.Sprintf("%s/_up", config.Get().CouchDB.Endpoint())
	res, err := http.Get(endpoint)
	if err != nil {
		return StatusUnhealthy
	}

	if res.StatusCode == http.StatusOK {
		return StatusHealthy
	}

	return StatusUnhealthy
}

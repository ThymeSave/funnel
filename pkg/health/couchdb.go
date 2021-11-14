package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thymesave/funnel/pkg/config"
)

// CouchDBCheck verifies couchdb is reachable
func CouchDBCheck(ctx context.Context) Status {
	endpoint := fmt.Sprintf("%s/_up", config.Get().CouchDB.Endpoint())
	res, err := http.Get(endpoint)
	if err != nil {
		return StatusUnhealthy
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return StatusHealthy
	}

	return StatusUnhealthy
}

package health

import "context"

// StatusHealthy is indicating a health-check succeeded
const StatusHealthy = "UP"

// StatusUnhealthy is indicating a health-check failed
const StatusUnhealthy = "DOWN"

// Status represents the health status of a given check
type Status string

// CheckFunc represents a health-check function
type CheckFunc func(ctx context.Context) Status

var checks = map[string]CheckFunc{}

// RegisterCheck registers a given health-check function under the given name
func RegisterCheck(name string, checkFunc CheckFunc) {
	checks[name] = checkFunc
}

// GetCheck by name or nil if none was found
func GetCheck(name string) *CheckFunc {
	check, ok := checks[name]
	if !ok {
		return nil
	}
	return &check
}

// GetAggregatedStatus checks all health checks and fails if any health-check reports unhealthy
func GetAggregatedStatus(ctx context.Context) Status {
	for _, check := range checks {
		if status := check(ctx); status != StatusHealthy {
			return StatusUnhealthy
		}
	}

	return StatusHealthy
}

func init() {
	RegisterCheck("general", GeneralCheck)
	RegisterCheck("couchdb", CouchDBCheck)
}

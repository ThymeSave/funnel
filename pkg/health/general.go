package health

import "context"

// GeneralCheck reports funnel status
func GeneralCheck(ctx context.Context) Status {
	return StatusHealthy
}

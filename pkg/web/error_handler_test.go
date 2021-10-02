package web

import (
	"net/http"
	"testing"

	"github.com/thymesave/funnel/pkg/config"
)

// TestNotFoundHandler runs against the real router and verifies not found is working
func TestNotFoundHandler(t *testing.T) {
	config.CreateDefault()
	rr := testHandler(CreateRouter(), "GET", "/404", nil)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %d", rr.Code)
	}
}

// TestMethodNotAllowedHandler runs against the single handler, because it is not possible to reliable test on
// the method not allowed.
func TestMethodNotAllowedHandler(t *testing.T) {
	rr := testHandler(http.HandlerFunc(MethodNotAllowedHandler), "GET", "/anywhere", nil)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, but got %d", rr.Code)
	}
}

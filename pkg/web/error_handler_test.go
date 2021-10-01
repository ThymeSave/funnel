package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNotFoundHandler runs against the real router and verifies not found is working
func TestNotFoundHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/testing/not-found", nil)
	rr := httptest.NewRecorder()
	CreateRouter().ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %d", rr.Code)
	}
}

// TestMethodNotAllowedHandler runs against the single handler, because it is not possible to reliable test on
// the method not allowed.
func TestMethodNotAllowedHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/testing/anywhere", nil)
	rr := httptest.NewRecorder()
	http.HandlerFunc(MethodNotAllowedHandler).ServeHTTP(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, but got %d", rr.Code)
	}
}

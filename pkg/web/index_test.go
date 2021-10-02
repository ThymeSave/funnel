package web

import "testing"

func TestIndexHandler(t *testing.T) {
	rr := testHandler(CreateRouter(), "GET", "/", nil)
	if rr.Code != 200 {
		t.Fatalf("Expected index to respond with 200, got %d instead", rr.Code)
	}
}

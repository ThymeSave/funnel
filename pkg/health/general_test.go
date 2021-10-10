package health

import (
	"context"
	"testing"
)

func TestGeneralCheck(t *testing.T) {
	if status := GeneralCheck(context.Background()); status != StatusHealthy {
		t.Fatal("Expected check to always be healthy")
	}
}

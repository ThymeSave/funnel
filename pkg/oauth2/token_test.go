package oauth2

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"testing"
)

func TestConvert(t *testing.T) {
	idToken := Convert(&oidc.IDToken{})
	if idToken == nil {
		t.Fatal("Expected convert to always succeed")
	}
}

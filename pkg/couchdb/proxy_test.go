package couchdb

import (
	"github.com/thymesave/funnel/pkg/oauth2"
	"net/http"
	"strconv"
	"testing"

	"github.com/thymesave/funnel/pkg/config"
)

func TestCreateModifyRequest(t *testing.T) {
	cfg := config.Get()
	cfg.Oauth2.UsernameClaim = "email"
	couchConfig := cfg.CouchDB

	modifier := CreateModifyRequest(cfg, "/couchdb")
	req, _ := http.NewRequest("GET", "/couchdb/test", nil)
	req.Header.Set("Cookie", "whatever")
	req.Header.Set("Authorization", "Bearer jwt-goes-here")
	modifier(req)

	if req.URL.Path != "/test" {
		t.Fatal("Prefix does not get removed")
	}

	if req.Header.Get("X-Auth-CouchDB-UserName") != "" {
		t.Fatal("X-Auth header should not be present when there is no JWT token sent")
	}

	if req.URL.Host != couchConfig.Host+":"+strconv.Itoa(couchConfig.Port) {
		t.Fatal("Target does not get overwritten")
	}

	if req.Header.Get("Authorization") != "" {
		t.Fatal("Authorization header is passed")
	}

	if req.Header.Get("Cookie") != "" {
		t.Fatal("Session header is passed")
	}

	if req.Header.Get("Content-Type") == "" {
		t.Fatal("Content-Type should be set")
	}

	// Set auth header
	idToken := oauth2.NewIDToken("https://auth.provider", "test")

	idToken.Claims = map[string]interface{}{
		"email": "test",
	}
	req = req.WithContext(oauth2.AddTokenToRequestContext(req, idToken))
	modifier(req)
	if req.Header.Get("X-Auth-CouchDB-UserName") == "" {
		t.Fatal("X-Auth header should be present when there is a JWT token sent")
	}
}

func TestModifyResponse(t *testing.T) {
	res := http.Response{
		Header: map[string][]string{
			"Server":              {"couchdb whatever"},
			"X-Couchdb-Body-Time": {"123r"},
			"Cookie":              {"crazy session stuff"},
		},
	}

	if err := ModifyResponse(&res); err != nil {
		t.Fatalf("Modifying response should NEVER fail, but failed with %s", err)
	}

	if res.Header.Get("Server") != "" {
		t.Fatalf("Server header should be removed in the response")
	}

	if res.Header.Get("Cookie") != "" {
		t.Fatalf("Cookie header should be removed in the response, as it might contain sensitive information")
	}
}

func TestCreateReverseProxy(t *testing.T) {
	proxy := CreateReverseProxy(config.Get(), "/couchdb")
	if proxy == nil {
		t.Fatal("CreateReverseProxy should NEVER be nil")
	}
}

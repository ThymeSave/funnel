package couchdb

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/thymesave/funnel/pkg/config"
)

func TestCreateModifyRequest(t *testing.T) {
	modifier := CreateModifyRequest("/couchdb")
	req, _ := http.NewRequest("GET", "/couchdb/test", nil)
	req.Header.Set("Cookie", "whatever")
	req.Header.Set("Authoriatzion", "Bearer jwt-goes-here")
	modifier(req)
	couchConfig := config.Get().CouchDB

	if req.URL.Path != "/test" {
		t.Fatal("Prefix does not get removed")
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
	proxy := CreateReverseProxy("/couchdb")
	if proxy == nil {
		t.Fatal("CreateReverseProxy should NEVER be nil")
	}
}

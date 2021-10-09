package web

import (
	"github.com/thymesave/funnel/pkg/config"
	"github.com/thymesave/funnel/pkg/oauth2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestSelfServiceSeedHandler(t *testing.T) {
	config.Get().Oauth2.UsernameClaim = "username"

	testCases := []struct {
		claims        map[string]interface{}
		handlerStatus int
		couchStatus   int
	}{
		{
			claims: map[string]interface{}{
				"username": "test",
			},
			handlerStatus: http.StatusCreated,
			couchStatus:   http.StatusCreated,
		},
		{
			claims:        map[string]interface{}{},
			handlerStatus: http.StatusInternalServerError,
			couchStatus:   http.StatusCreated,
		},
		{
			claims: map[string]interface{}{
				"username": "test",
			},
			handlerStatus: http.StatusOK,
			couchStatus:   http.StatusConflict,
		},
		{
			claims: map[string]interface{}{
				"username": "test",
			},
			handlerStatus: http.StatusInternalServerError,
			couchStatus:   http.StatusBadGateway,
		},
	}

	for _, tc := range testCases {
		couchMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(tc.couchStatus)
		}))
		couchURL, _ := url.Parse(couchMock.URL)
		port, _ := strconv.Atoi(couchURL.Port())
		config.Get().CouchDB = &config.CouchDB{
			Scheme:    couchURL.Scheme,
			Host:      strings.Split(couchURL.Host, ":")[0],
			Port:      port,
			AdminUser: "admin",
		}

		token := oauth2.NewIDToken("https://auth.provider", "test")
		token.Claims = tc.claims
		r, _ := http.NewRequest("GET", "/foo", nil)
		r = r.WithContext(oauth2.AddTokenToRequestContext(r, token))

		rr := testHandlerWithRequest(http.HandlerFunc(SelfServiceSeedHandler), r)

		if rr.Code != tc.handlerStatus {
			t.Fatalf("Expected http status %d, but got %d", tc.handlerStatus, rr.Code)
		}

		couchMock.Close()
	}
}

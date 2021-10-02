package web

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/thymesave/funnel/pkg/config"
)

func testHandler(handler http.Handler, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func init() {
	// Create default version for tests
	config.CreateDefault()

	// Init couchdb proxy
	CreateCouchDBReverseProxy()
}

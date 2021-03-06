package web

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/thymesave/funnel/pkg/config"
)

func mockResponseWithContentType(httpStatus int, responseBody string, contentType string, testFunc func(*httptest.Server)) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(httpStatus)
		_, _ = w.Write([]byte(responseBody))
	}))
	testFunc(server)
	defer server.Close()
}

func mockResponse(httpStatus int, responseBody string, testFunc func(*httptest.Server)) {
	mockResponseWithContentType(httpStatus, responseBody, "text/plain", testFunc)
}

func testHandlerWithRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func testHandler(handler http.Handler, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	return testHandlerWithRequest(handler, req)
}

func createRouter() http.Handler {
	_ = patchOidConfig("oidc-config")
	r, err := CreateRouter(context.Background())
	if err != nil {
		println(err.Error())
	}
	return r
}

func patchOidConfig(configPath string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// Set to current package and testdata folder
	config.Get().Oauth2.IssuerURL = "file://" + path + "/testdata/" + configPath
	return nil
}

func init() {
	os.Setenv("FUNNEL_OAUTH2_ISSUER_URL", "")
	os.Setenv("FUNNEL_OAUTH2_CLIENT_ID", "")

	// Create default version for tests
	config.CreateDefault()

	// Init couchdb proxy
	CreateCouchDBReverseProxy(config.Get())
}

func TestCreateRouter(t *testing.T) {
	_ = patchOidConfig("oidc-config")
	_, err := CreateRouter(context.Background())
	if err != nil {
		t.Fatalf("Expected router creation to be successful, but got %s", err.Error())
	}

	_ = patchOidConfig("invalid-oidc-config")
	_, err = CreateRouter(context.Background())
	if err == nil {
		t.Fatal("Expected router creation to be fail with invalid oidc config")
	}
}

package web

import (
	"net/http"
	"testing"
)

func TestCouchDbProxyHandler(t *testing.T) {
	testCases := []struct {
		path           string
		expectedStatus int
	}{
		{
			"/",
			CouchDBProxyHiddenPathStatus,
		},
		{
			"/_utils",
			CouchDBProxyHiddenPathStatus,
		},
		{
			"/_utils/what/ever",
			CouchDBProxyHiddenPathStatus,
		},
		{
			"/some/legit/url/someone/might/access",
			http.StatusOK,
		},
	}

	// Overwrite for testing purpose to ensure whether couchdb is running or not the test is not different
	couchdbReverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusOK)
	}
	couchdbReverseProxy.ModifyResponse = func(response *http.Response) error {
		response.StatusCode = 200
		return nil
	}

	for _, tc := range testCases {
		rr := testHandler(http.HandlerFunc(CouchDbProxyHandler), "GET", PathCouchDbService+tc.path, nil)
		if rr.Code != tc.expectedStatus {
			t.Fatalf("Expected status code %d, but got %d for path %s", tc.expectedStatus, rr.Code, tc.path)
		}
	}
}

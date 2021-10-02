package web

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/thymesave/funnel/pkg/couchdb"
)

var couchdbReverseProxy *httputil.ReverseProxy
var couchDbHiddenPaths = []string{
	// Index contains sensitive information
	"/",
	// Possible leak of passwords - misuse
	"/_session",
}

var couchDbHiddenPrefixes = []string{
	// Admin interface
	"/_utils",
}

// CouchDBProxyHiddenPathStatus is the http status that gets delivered to users in case the underlying
// endpoint should not be reachable via reverse proxy
const CouchDBProxyHiddenPathStatus = http.StatusTeapot

// CreateCouchDBReverseProxy must be called after configuration has been loaded and before executing
//requests against the couchdb services endpoint(s)
func CreateCouchDBReverseProxy() {
	couchdbReverseProxy = couchdb.CreateReverseProxy(PathCouchDbService)
}

// CouchDbProxyHandler is the proxy for couchdb
func CouchDbProxyHandler(w http.ResponseWriter, r *http.Request) {
	rPath := strings.TrimPrefix(r.URL.Path, PathCouchDbService)

	// Check for direct paths that should not be delivered
	for _, path := range couchDbHiddenPaths {
		if rPath == path {
			w.WriteHeader(CouchDBProxyHiddenPathStatus)
			return
		}
	}

	// Check for path prefixes that should not be delivered
	for _, path := range couchDbHiddenPrefixes {
		if strings.HasPrefix(rPath, path) {
			w.WriteHeader(CouchDBProxyHiddenPathStatus)
			return
		}
	}

	couchdbReverseProxy.ServeHTTP(w, r)
}

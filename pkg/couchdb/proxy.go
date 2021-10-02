package couchdb

import (
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/thymesave/funnel/pkg/config"
)

// GetCouchDBUser from the request
func GetCouchDBUser(r *http.Request) string {
	// TODO Set based on JWT
	return "admin"
}

// CreateModifyRequest returns a function capable of modifying requests to make requests to couchdb without copying
// all headers
func CreateModifyRequest(basePath string) func(r *http.Request) {
	couchdbConfig := config.Get().CouchDB
	host := couchdbConfig.Host + ":" + strconv.Itoa(couchdbConfig.Port)
	scheme := couchdbConfig.Scheme

	return func(r *http.Request) {
		header := http.Header{}
		header.Set("X-Auth-CouchDB-UserName", GetCouchDBUser(r))
		header.Set("X-Auth-CouchDB-Roles", "")
		header.Set("Content-Type", "application/json")
		header.Set("Accept", "application/json")

		r.URL.Scheme = scheme
		r.URL.Host = host
		r.URL.Path = strings.TrimPrefix(r.URL.Path, basePath)
		r.Header = header
	}
}

// ModifyResponse and remove sensitive details
func ModifyResponse(r *http.Response) error {
	// Hide server versions and couchdb specific headers
	r.Header.Del("Server")
	r.Header.Del("X-Couchdb-Body-Time")
	r.Header.Del("Cookie")
	return nil
}

// CreateReverseProxy for CouchDB
func CreateReverseProxy(basePath string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director:       CreateModifyRequest(basePath),
		ModifyResponse: ModifyResponse,
	}
}

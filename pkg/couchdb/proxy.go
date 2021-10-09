package couchdb

import (
	"github.com/thymesave/funnel/pkg/oauth2"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/thymesave/funnel/pkg/config"
)

// GetCouchDBUser from the request
func GetCouchDBUser(r *http.Request, oauth2Config *config.OAuth2) (string, error) {
	return oauth2.GetUsername(r, oauth2Config)
}

// CreateModifyRequest returns a function capable of modifying requests to make requests to couchdb without copying
// all headers
func CreateModifyRequest(cfg *config.AppConfig, basePath string) func(r *http.Request) {
	couchdbConfig := cfg.CouchDB
	oauth2Config := cfg.Oauth2
	host := couchdbConfig.Host + ":" + strconv.Itoa(couchdbConfig.Port)
	scheme := couchdbConfig.Scheme

	return func(r *http.Request) {
		header := http.Header{}

		username, err := GetCouchDBUser(r, oauth2Config)
		if err == nil {
			setAuthHeader(&header, username, []string{})
		}

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
func CreateReverseProxy(cfg *config.AppConfig, basePath string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director:       CreateModifyRequest(cfg, basePath),
		ModifyResponse: ModifyResponse,
	}
}

package couchdb

import (
	"net/http"
	"strings"
)

func setAuthHeader(h *http.Header, username string, roles []string) {
	h.Set("X-Auth-CouchDB-UserName", username)
	h.Set("X-Auth-CouchDB-Roles", strings.Join(roles, ","))
}

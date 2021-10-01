package web

import "net/http"

// NotFoundHandler sets the HTTP status and finishes
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

// MethodNotAllowedHandler sets the HTTP status and finishes
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

package web

import "net/http"

// IndexHandler is the default response a curious user will see
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Return useful information
	_ = SendJSON(w, map[string]interface{}{
		"status": "up",
	})
}

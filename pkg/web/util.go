package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendJSON response for the given response writer, if an error occurs
// and 500 is returned to the client and the error is given back if further
//error handling is required. The cause for the error is also logged.
func SendJSON(w http.ResponseWriter, payload interface{}) error {
	serialized, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to serialize json: %e", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(serialized)
	return nil
}

package web

import (
	"github.com/thymesave/funnel/pkg/config"
	"github.com/thymesave/funnel/pkg/couchdb"
	"github.com/thymesave/funnel/pkg/oauth2"
	"log"
	"net/http"
)

// SelfServiceSeedHandler creates initial couchdb resources if not already present
func SelfServiceSeedHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Get()
	username, err := oauth2.GetUsername(r, cfg.Oauth2)
	if err != nil {
		log.Printf("Failed to get username from request: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := couchdb.CreateUser(couchdb.NewClient(), cfg.CouchDB, username)
	if err != nil {
		log.Println("Failed to create user: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpStatus := http.StatusOK

	if res.NewUser {
		httpStatus = http.StatusCreated
	}

	_ = SendJSON(w, httpStatus, res)
}

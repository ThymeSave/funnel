package web

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)
import "github.com/gorilla/mux"

func addMiddlewares(r *mux.Router) http.Handler {
	loggingHandler := handlers.LoggingHandler(os.Stdout, r)
	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowCredentials(), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}))(loggingHandler)
	return corsHandler
}

// CreateRouter returns a ready to use router
func CreateRouter() http.Handler {
	r := mux.NewRouter()
	registerMetricHandler(r)
	return addMiddlewares(r)
}

package web

import (
	"net/http"
	"os"

	"github.com/thymesave/funnel/pkg/config"

	"github.com/gorilla/handlers"
)
import "github.com/gorilla/mux"

func addMiddlewares(r *mux.Router) http.Handler {
	webConfig := config.Get().Web
	loggingHandler := handlers.LoggingHandler(os.Stdout, r)
	corsHandler := handlers.CORS(handlers.AllowedOrigins(webConfig.CORSOrigins), handlers.AllowCredentials(), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}))(loggingHandler)
	return corsHandler
}

func createHandler(r *mux.Router, handlerFunc http.HandlerFunc) http.Handler {
	return r.NewRoute().HandlerFunc(handlerFunc).GetHandler()
}

func registerAppRoutes(r *mux.Router) {
	r.Path("/").Methods("GET").HandlerFunc(IndexHandler)
}

// CreateRouter returns a ready to use router
func CreateRouter() http.Handler {
	r := mux.NewRouter()
	registerAppRoutes(r)
	r.NotFoundHandler = createHandler(r, NotFoundHandler)
	r.MethodNotAllowedHandler = createHandler(r, MethodNotAllowedHandler)
	registerMetricHandler(r)
	return addMiddlewares(r)
}

package web

import (
	"context"
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

func registerAppRoutes(ctx context.Context, r *mux.Router) error {
	oauth2Middleware, err := CreateOAuth2Handler(ctx)
	if err != nil {
		return err
	}

	r.Path("/").Methods("GET").HandlerFunc(IndexHandler)

	r.PathPrefix(PathCouchDbService + "/").HandlerFunc(oauth2Middleware(CouchDbProxyHandler))
	return nil
}

// CreateRouter returns a ready to use router
func CreateRouter(ctx context.Context) (http.Handler, error) {
	r := mux.NewRouter()
	err := registerAppRoutes(ctx, r)
	if err != nil {
		return nil, err
	}
	registerMetricHandler(r)
	r.NotFoundHandler = createHandler(r, NotFoundHandler)
	r.MethodNotAllowedHandler = createHandler(r, MethodNotAllowedHandler)
	return addMiddlewares(r), nil
}

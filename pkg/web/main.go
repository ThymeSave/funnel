package web

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"

	"github.com/thymesave/funnel/pkg/config"

	"github.com/gorilla/handlers"
)
import "github.com/gorilla/mux"

func addGlobalMiddlewares(r *mux.Router) http.Handler {
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

	r.Path("/").Methods(http.MethodGet).HandlerFunc(IndexHandler)
	r.Path("/health").Methods(http.MethodGet).HandlerFunc(HealthHandler)
	r.Path("/health/{component}").Methods(http.MethodGet).HandlerFunc(HealthHandler)
	r.Path("/metrics").Methods(http.MethodGet).Handler(promhttp.Handler())
	r.Path("/self-service/db").Methods(http.MethodPut).HandlerFunc(oauth2Middleware(SelfServiceSeedHandler))
	r.PathPrefix(PathCORSProxy + "/").Methods(http.MethodGet).HandlerFunc(oauth2Middleware(CORSProxyHandler))
	r.PathPrefix(PathCouchDbService + "/").HandlerFunc(oauth2Middleware(CouchDbProxyHandler))

	return nil
}

// CreateRouter returns a ready to use router
func CreateRouter(ctx context.Context) (http.Handler, error) {
	r := mux.NewRouter().UseEncodedPath()

	err := registerAppRoutes(ctx, r)
	if err != nil {
		return nil, err
	}

	r.NotFoundHandler = createHandler(r, NotFoundHandler)
	r.MethodNotAllowedHandler = createHandler(r, MethodNotAllowedHandler)

	return addGlobalMiddlewares(r), nil
}

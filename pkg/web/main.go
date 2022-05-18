package web

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thymesave/funnel/pkg/config"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
)
import "github.com/gorilla/mux"

var (
	allowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders = []string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "Origin", "Authorization", "User-Agent", "Referer"}
)

func addGlobalMiddlewares(r *mux.Router) http.Handler {
	webConfig := config.Get().Web
	loggingHandler := handlers.LoggingHandler(os.Stdout, r)
	corsHandler := handlers.CORS(handlers.AllowedOrigins(webConfig.CORSOrigins), handlers.AllowedHeaders(allowedHeaders), handlers.IgnoreOptions(), handlers.AllowCredentials(), handlers.AllowedMethods(allowedMethods))(loggingHandler)
	return corsHandler
}

func createHandler(r *mux.Router, handlerFunc http.HandlerFunc) http.Handler {
	return r.NewRoute().HandlerFunc(handlerFunc).GetHandler()
}

func createCorsHandler(webConfig *config.HTTP) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(webConfig.CORSOrigins, ","))
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ","))
	}
}

func registerAppRoutes(ctx context.Context, r *mux.Router) error {
	oauth2Middleware, err := CreateOAuth2Handler(ctx)
	if err != nil {
		return err
	}

	webConfig := config.Get().Web
	corsHandler := createCorsHandler(webConfig)

	r.Path("/").Methods(http.MethodGet).HandlerFunc(IndexHandler)
	r.Path("/health").Methods(http.MethodGet).HandlerFunc(HealthHandler)
	r.Path("/health/{component}").Methods(http.MethodGet).HandlerFunc(HealthHandler)
	r.Path("/metrics").Methods(http.MethodGet).Handler(promhttp.Handler())
	r.Path("/self-service/db").Methods(http.MethodPut).HandlerFunc(oauth2Middleware(SelfServiceSeedHandler))
	r.PathPrefix(PathCORSProxy + "/").Methods(http.MethodGet).HandlerFunc(CORSProxyHandler)
	r.PathPrefix(PathCouchDbService + "/").Methods(http.MethodOptions).HandlerFunc(corsHandler)
	r.PathPrefix(PathCouchDbService + "/").HandlerFunc(oauth2Middleware(CouchDbProxyHandler))
	r.Methods(http.MethodOptions).HandlerFunc(corsHandler)

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

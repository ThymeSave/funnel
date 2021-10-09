package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/thymesave/funnel/pkg/config"
	"github.com/thymesave/funnel/pkg/web"
)

// RunHTTPServer starts the http server
func RunHTTPServer(ctx context.Context, port int) error {
	listen := "0.0.0.0:" + strconv.Itoa(port)
	log.Println("Starting funnel on  " + listen)

	router, err := web.CreateRouter(ctx)
	if err != nil {
		return err
	}

	return http.ListenAndServe(listen, router)
}

func terminate(status int) {
	os.Exit(status)
}

// Start funnel
func Start() {
	appCtx := context.Background()

	if err := config.ReadConfig(appCtx); err != nil {
		log.Printf("Failed to read config: %s", err)
		terminate(2)
	}

	web.CreateCouchDBReverseProxy(config.Get())

	if err := RunHTTPServer(appCtx, config.Get().Web.Port); err != nil {
		log.Printf("HTTP-Server crashed: %s", err)
		terminate(1)
	}

	terminate(0)
}

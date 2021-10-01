package cmd

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thymesave/funnel/pkg/web"
)

// RunHTTPServer starts the http server
func RunHTTPServer(port int) error {
	listen := "0.0.0.0:" + strconv.Itoa(port)
	log.Println("Starting server on  " + listen)
	return http.ListenAndServe(listen, web.CreateRouter())
}

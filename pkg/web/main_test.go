package web

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func testHandler(handler http.Handler, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

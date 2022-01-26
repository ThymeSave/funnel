package web

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/thymesave/funnel/pkg/buildinfo"
	"io"
	"net/http"
	"strings"
	"time"
)

// StatusRequestFailed indicates the request to the origin has failed
const StatusRequestFailed = "ORIGIN_REQUEST_FAILED"

// StatusReadFailed indicates during reading from the origin a problem occurred
const StatusReadFailed = "ORIGIN_RESPONSE_READ_FAILED"

// StatusInvalidContentType indicates that the response content type is not supported by funnel
const StatusInvalidContentType = "ORIGIN_RESPONSE_CONTENT_TYPE_UNSUPPORTED"

type proxyError struct {
	Status           string `json:"errorStatus"`
	UpstreamResponse string `json:"upstreamResponse"`
}

var (
	proxyTransferredBytesMetric = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "cors_proxy_transferred_bytes",
		Help: "Transferred bytes via CORS Proxy",
	})
	proxyHTTPClient = http.Client{
		Timeout: 3 * time.Second,
	}
)

// CORSProxyHandler contains functionality to route GET request to other origins
func CORSProxyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "ThymeSave_Funnel/"+buildinfo.Version)
	res, err := proxyHTTPClient.Do(req)
	if err != nil {
		_ = SendJSON(w, http.StatusBadGateway, proxyError{StatusRequestFailed, ""})
		return
	}

	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		_ = SendJSON(w, http.StatusBadRequest, proxyError{StatusInvalidContentType, ""})
		return
	}

	defer res.Body.Close()
	w.Header().Set("Content-Type", "text/plain")
	written, err := io.Copy(w, res.Body)
	if err != nil {
		_ = SendJSON(w, http.StatusInternalServerError, proxyError{StatusReadFailed, ""})
		return
	}

	proxyTransferredBytesMetric.Observe(float64(written))
}

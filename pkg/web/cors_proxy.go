package web

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/thymesave/funnel/pkg/buildinfo"
	"github.com/thymesave/funnel/pkg/util"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// StatusRequestFailed indicates the request to the origin has failed
const StatusRequestFailed = "ORIGIN_REQUEST_FAILED"

// StatusReadFailed indicates during reading from the origin a problem occurred
const StatusReadFailed = "ORIGIN_RESPONSE_READ_FAILED"

// StatusInvalidContentType indicates that the response content type is not supported by funnel
const StatusInvalidContentType = "ORIGIN_RESPONSE_CONTENT_TYPE_UNSUPPORTED"

// StatusInvalidURL indicates that the requested url is not valid and can not be requested
const StatusInvalidURL = "INVALID_URL"

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

func isForceAllowLocal() bool {
	return os.Getenv("FUNNEL_CORS_FORCE_LOCAL_RESOLUTION") == "true"
}

func newProxyError(status string, upstreamResponse string) proxyError {
	return proxyError{status, ""}
}

// CORSProxyHandler contains functionality to route GET request to other origins
func CORSProxyHandler(w http.ResponseWriter, r *http.Request) {
	corsURL := r.URL.Query().Get("url")
	if corsURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedURL, err := url.Parse(corsURL)
	if err != nil {
		_ = SendJSON(w, http.StatusBadRequest, newProxyError(StatusInvalidURL, ""))
		return
	}

	if !isForceAllowLocal() && util.ResolvesHostnameToLocalIP(parsedURL.Host) {
		_ = SendJSON(w, http.StatusBadRequest, newProxyError(StatusInvalidURL, ""))
	}

	req, _ := http.NewRequest("GET", corsURL, nil)
	req.Header.Set("User-Agent", "ThymeSave_Funnel/"+buildinfo.Version)
	res, err := proxyHTTPClient.Do(req)
	if err != nil {
		_ = SendJSON(w, http.StatusBadGateway, newProxyError(StatusRequestFailed, ""))
		return
	}

	if !strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		_ = SendJSON(w, http.StatusBadRequest, newProxyError(StatusInvalidContentType, ""))
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

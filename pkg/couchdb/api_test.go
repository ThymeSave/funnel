package couchdb

import (
	"github.com/thymesave/funnel/pkg/config"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func mockResponse(httpStatus int, responseBody string, testFunc func(*httptest.Server, Client)) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(httpStatus)
		_, _ = w.Write([]byte(responseBody))
	}))
	testFunc(server, Client{HTTPClient: server.Client()})
	defer server.Close()
}

func TestCreateUserDB(t *testing.T) {
	testCases := []struct {
		httpStatus   int
		responseBody string
		newUser      bool
		errMsg       string
	}{
		{
			httpStatus:   http.StatusCreated,
			responseBody: "{}",
			newUser:      true,
		},
		{
			httpStatus:   http.StatusConflict,
			responseBody: "{}",
			newUser:      false,
		},
		{
			httpStatus:   http.StatusInternalServerError,
			responseBody: "{}",
			newUser:      false,
			errMsg:       "failed to create user, upstream response: {}",
		},
	}

	for _, tc := range testCases {
		mockResponse(tc.httpStatus, tc.responseBody, func(srv *httptest.Server, c Client) {
			testServerURL, _ := url.Parse(srv.URL)
			port, _ := strconv.Atoi(testServerURL.Port())
			res, err := CreateUser(&c, &config.CouchDB{
				Scheme:    testServerURL.Scheme,
				Host:      strings.Split(testServerURL.Host, ":")[0],
				Port:      port,
				AdminUser: "admin",
			}, "test")
			if err != nil {
				if tc.errMsg == "" {
					t.Fatal(err)
				} else if err.Error() != tc.errMsg {
					t.Fatalf("Expected errror message %s got %s", tc.errMsg, err)
				}
				return
			}

			if tc.newUser != res.NewUser {
				t.Fatalf("Expected newUser to be %t but got %t", tc.newUser, res.NewUser)
			}
		})
	}

	_, err := CreateUser(&Client{HTTPClient: http.DefaultClient}, &config.CouchDB{Scheme: "https", AdminUser: "admin"}, "t")
	if err == nil {
		t.Fatal("Expected to return error in case of invalid config")
	}
}

func TestNewClient(t *testing.T) {
	_ = NewClient()
}

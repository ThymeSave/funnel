package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSON(t *testing.T) {
	testCases := []struct {
		expectedBody string
		input        interface{}
	}{
		{
			expectedBody: "{\"foo\":\"bar\"}",
			input:        map[string]string{"foo": "bar"},
		},
		{
			expectedBody: "",
			input:        make(chan int),
		},
	}

	for _, tc := range testCases {
		dummyHandler := func(w http.ResponseWriter, r *http.Request) {
			_ = SendJSON(w, http.StatusOK, tc.input)
		}
		req, _ := http.NewRequest("GET", "/any-resource", nil)
		rr := httptest.NewRecorder()
		http.HandlerFunc(dummyHandler).ServeHTTP(rr, req)
		res, _ := ioutil.ReadAll(rr.Body)
		if string(res) != tc.expectedBody && (tc.expectedBody != "") {
			t.Errorf("Expected json to be %s, but got %s", tc.expectedBody, string(res))
		}
	}

}

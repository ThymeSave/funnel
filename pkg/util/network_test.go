package util

import (
	"os"
	"testing"
)

func TestResolvesHostnameToLocalIP(t *testing.T) {
	hostname, _ := os.Hostname()
	testCases := []struct {
		hostname       string
		expectedResult bool
	}{
		{
			hostname:       "localhost",
			expectedResult: true,
		},
		{
			hostname:       "google.com",
			expectedResult: false,
		},
		{
			hostname:       "foo.bar.what.ever",
			expectedResult: false,
		},
		{
			hostname:       hostname,
			expectedResult: true,
		},
	}
	for _, tc := range testCases {
		if ResolvesHostnameToLocalIP(tc.hostname) != tc.expectedResult {
			t.Fatalf("Expected hostname %s to resolve %v", tc.hostname, tc.expectedResult)
		}
	}
}

package util

import (
	"net"
)

// ResolvesHostnameToLocalIP checks if the given hostname resolves to a logback or private ip address
func ResolvesHostnameToLocalIP(hostname string) bool {
	if hostname == "localhost" {
		return true
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		return false
	}

	for _, hostIP := range ips {
		if hostIP.IsLoopback() || hostIP.IsPrivate() {
			return true
		}
	}

	return false
}

package VAS_util_go

import "net"

// Extracts the IP from a string containing both port and IP
func ExtractIP(remote string) string {
	host, _, err := net.SplitHostPort(remote)
	if err != nil {
		return ""
	}
	return host
}

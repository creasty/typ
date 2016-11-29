package string_h

import (
	"strings"
)

func IsLocalHostname(host string) bool {
	isLocal := strings.Index(host, "localhost") == 0 ||
		strings.Index(host, "dockerhost") == 0 ||
		strings.Index(host, "test") == 0 ||
		strings.Index(host, "127.0.") == 0 ||
		strings.Index(host, "192.168.") == 0 ||
		strings.Index(host, "10.") == 0 ||
		strings.Index(host, "176.") == 0

	return isLocal
}

package docfmt

import (
	"net"
	"regexp"
	"strings"
)

const (
	DocumentationIPv4       = "192.0.2.1"
	DocumentationIPv4Prefix = "192.0.2.0/24"
)

var ipv4LiteralPattern = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}(?:/[0-9]{1,2})?\b`)

// CanonicalizeNetworkLiterals replaces globally routable IPv4 examples with
// RFC 5737 documentation addresses. Private, local, invalid, and already-safe
// documentation addresses are preserved.
func CanonicalizeNetworkLiterals(content string) string {
	return ipv4LiteralPattern.ReplaceAllStringFunc(content, func(match string) string {
		address := match
		hasPrefix := false
		if slash := strings.IndexByte(match, '/'); slash >= 0 {
			address = match[:slash]
			hasPrefix = true
		}

		ip := net.ParseIP(address)
		if ip == nil || ip.To4() == nil || !ip.IsGlobalUnicast() || ip.IsPrivate() || isDocumentationIPv4(ip) {
			return match
		}
		if hasPrefix {
			return DocumentationIPv4Prefix
		}
		return DocumentationIPv4
	})
}

func isDocumentationIPv4(ip net.IP) bool {
	v4 := ip.To4()
	if v4 == nil {
		return false
	}
	return (v4[0] == 192 && v4[1] == 0 && v4[2] == 2) ||
		(v4[0] == 198 && v4[1] == 51 && v4[2] == 100) ||
		(v4[0] == 203 && v4[1] == 0 && v4[2] == 113)
}

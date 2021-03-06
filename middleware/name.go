package middleware

import (
	"strings"

	"github.com/miekg/dns"
)

// Name represents a domain name.
type Name string

// Matches checks to see if other matches n.
//
// Name matching will probably not always be a direct
// comparison; this method assures that names can be
// easily and consistently matched.
func (n Name) Matches(other string) bool {
	return strings.HasSuffix(string(n), other)
}

// Normalize lowercases and makes n fully qualified.
func (n Name) Normalize() string {
	return strings.ToLower(dns.Fqdn(string(n)))
}

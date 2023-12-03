package auth

import (
	"net/http"
	"strings"
)

// ej localhost:8080 to localhost
func (a *Auth) setDomain(r *http.Request) {
	if a.domain == "" {
		hostParts := strings.Split(r.Host, ":")
		a.domain = hostParts[0]
	}
}

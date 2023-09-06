package auth

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
)

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.1.html
// size ej: 16, 32, 64
func buildUniqueKey(size_key uint8) string {
	b := make([]byte, size_key)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ej localhost:8080 to localhost
func (a *Auth) setDomain(r *http.Request) {
	if a.domain == "" {
		hostParts := strings.Split(r.Host, ":")
		a.domain = hostParts[0]
	}
}

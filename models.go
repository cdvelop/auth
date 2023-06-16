package auth

import (
	"net/http"
)

type provider interface {
	Name() string
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
}

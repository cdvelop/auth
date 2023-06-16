package auth

import (
	"net/http"
)

func (a auth) addHttpMuxHandlers(mux *http.ServeMux) {

	for _, p := range a.providers {
		mux.HandleFunc("/login/"+p.Name(), p.Login)
		mux.HandleFunc("/callback/"+p.Name(), p.Callback)
	}

}

package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var client_id_google string     //go build -ldflags "-X auth.client_id_google=XXX"
var client_secret_google string //go build -ldflags "-X auth.client_secret_google=XXX"

type google_dev struct {
	*auth
	*oauth2.Config
}

// proveedor google cuenta desarrollador
func (google_dev) Name() string {
	return "google_dev"
}

func ConfigGoogleOauth2() *oauth2.Config {

	if client_id_google == "" {
		client_id_google = os.Getenv("client_id_google")
		if client_id_google == "" {
			showErrorAndExit("variable vacía client_id_google")
		}
	}

	if client_secret_google == "" {
		client_secret_google = os.Getenv("client_secret_google")
		if client_secret_google == "" {
			showErrorAndExit("variable vacía client_secret_google")
		}
	}

	return &oauth2.Config{
		ClientID:     client_id_google,
		ClientSecret: client_secret_google,
		Endpoint:     google.Endpoint,
		RedirectURL:  "",
		Scopes:       []string{},
	}
}

//{"web":{"client_id":"XXX","project_id":"XXX","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"XXXX","redirect_uris":["http://localhost:8080/callback"]}}

func (p google_dev) Login(w http.ResponseWriter, r *http.Request) {

	_, err := p.getTokenFromClientCookie(r)
	if err != nil { //no hay token solicitar uno nuevo

		state := p.createStatusCookie(w)

		authURL := p.Config.AuthCodeURL(state, oauth2.AccessTypeOffline)

		http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, p.home, http.StatusSeeOther)
	}

}

func (p google_dev) Callback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Recorre los parámetros enviados en la URL")
	for key, values := range r.Form {
		fmt.Printf("%s: %s\n", key, values)
	}
	fmt.Println("---------------------------")

	receivedState := r.Form.Get(p.status)

	cookie_state, err := getCookie(p.status, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if receivedState != cookie_state.Value {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")

	token, err := p.Config.Exchange(context.TODO(), code)
	if err != nil {
		http.Error(w, "Error Exchange code for token "+err.Error(), http.StatusInternalServerError)
		return
	}

	p.saveTokenInClientCookie(token, w)

	http.Redirect(w, r, p.home, http.StatusSeeOther)

}

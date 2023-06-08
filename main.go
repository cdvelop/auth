package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	clientID     = "your-client-id"
	clientSecret = "your-client-secret"
	redirectURL  = "http://localhost:8080/callback"
	authURL      = "http://localhost:8080/auth"
	tokenURL     = "http://localhost:8080/token"
)

var (
	oauthConfig *oauth2.Config
	state       string
)

func main() {
	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"read", "write"}, // Permisos solicitados por la aplicación
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, OAuth2!")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state = "random-state-string" // Genera un estado aleatorio para proteger contra ataques CSRF

	authURL := oauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	receivedState := r.URL.Query().Get("state")

	if receivedState != state {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return
	}

	// Utiliza el token para acceder a los recursos protegidos
	// Aquí puedes realizar solicitudes HTTP con el token de acceso a tus APIs internas

	fmt.Fprint(w, "Authentication successful!")
}

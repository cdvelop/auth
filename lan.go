package auth

import (
	"fmt"
	"net/http"
)

// proveedor red area local
type lan struct {
	*Auth
	redirect_data string
}

func (p lan) LoginPattern() string {
	return "/login/lan"
}

// redirect_data ej: "/home", "/"
func (a *Auth) UseLanAuth(redirect_data string) *lan {

	p := lan{
		Auth:          a,
		redirect_data: redirect_data,
	}

	if a.mux != nil {
		a.mux.HandleFunc(p.LoginPattern(), p.login)
	}

	return &p
}

// Lan login
func (l *lan) login(w http.ResponseWriter, r *http.Request) {
	l.setDomain(r)

	if r.Method != http.MethodPost {
		http.Error(w, "Método http "+r.Method+" no permitido", http.StatusNotAcceptable)
		return
	}

	r.ParseForm()
	login_data := map[string]string{}
	for key, values := range r.Form {
		if len(values) == 1 {
			login_data[key] = values[0]
		}
	}

	err := l.object.ValidateData(false, false, login_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	client_data := l.ReadObject(l.client_table_name, login_data)
	if len(client_data) == 0 {
		http.Error(w, "Datos de Session Incorrectos", http.StatusNotAcceptable)
		return
	}

	fmt.Println("DATA CLIENTE: ", client_data)

	l.createStatusCookie(w)

	fmt.Println("REDIRECT URL: ", l.redirect_data)

	// time.Sleep(400 * time.Millisecond)

	// http.Redirect(w, r, l.redirect_data, http.StatusTemporaryRedirect)
	http.Redirect(w, r, l.redirect_data, http.StatusOK)

}

package auth_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cdvelop/auth"
)

func login(w http.ResponseWriter, r *http.Request) {
	auth.SetCookie("session", "example", "", false, 2*time.Minute, w)
}

func home(w http.ResponseWriter, r *http.Request) {

	// Leer la cookie aquí
	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value != "example" {
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized)
		return
	}
	// Resto de la lógica del controlador home

	fmt.Println("COOKIE: ", cookie)

}

func TestCookieCreationAndReading(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/home", home)

	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	// Realizar solicitud al endpoint "/login" para crear la cookie
	loginURL := server.URL + "/login"
	resp, err := client.Get(loginURL)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	// Verificar si se creó la cookie
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		t.Errorf("La cookie no fue creada")
	}

	// Crear una nueva solicitud para leer la cookie
	homeURL := server.URL + "/home"
	req, err := http.NewRequest("GET", homeURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Agregar la cookie a la solicitud
	req.AddCookie(cookies[0])

	// Enviar la solicitud a la función home
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	// Verificar si la respuesta fue exitosa (código 200)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("El acceso fue denegado. Código de estado: %d", resp.StatusCode)
	}
}

package auth_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cdvelop/auth"
)

// var receivedCookies []*http.Cookie

func TestLanLoginDefault(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/", db{}.Home)

	a := auth.Add("/", "clients", db{}, false, mux)

	auth_lan := a.UseLanAuth("/")

	server := httptest.NewServer(mux)
	defer server.Close()

	for prueba, data := range testData {
		t.Run((prueba), func(t *testing.T) {

			client := &http.Client{}

			resp, err := client.PostForm(server.URL+auth_lan.LoginPattern(), data.form_values)
			// resp, err := http.PostForm(server.URL+auth_lan.LoginPattern(), data.form_values)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			// Verificar si se creó la cookie
			cookies := resp.Cookies()
			if len(cookies) == 0 {
				t.Errorf("La cookie no fue creada")
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if resp.Status != data.expected {
				fmt.Println("RESPUESTA: ", string(bodyBytes))
				fmt.Println("CÓDIGO: ", resp.Status, "SE ESPERABA: ", data.expected)
				log.Fatalln()
			}

			// Crear una nueva solicitud para leer la cookie
			req, err := http.NewRequest("GET", server.URL+"/", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Agregar la cookie a la solicitud
			if len(cookies) != 0 {
				req.AddCookie(cookies[0])
			}

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

		})
	}

}

type db struct{}

func (db) Home(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("Lan Login Recorre los parámetros enviados en la URL")
	login_data := map[string]string{}
	for key, values := range r.Form {
		if len(values) == 1 {
			login_data[key] = values[0]
		}
		fmt.Printf("%s: %s\n", key, values)
	}
	fmt.Println("---------------------------")

	// Leer cookies
	for _, c := range r.Cookies() {

		fmt.Println("Name: ", c.Name, " value: ", c.Value)

	}
}

// Función auxiliar para simular la obtención de objetos de la base de datos
func getObjectsFromDB() []map[string]string {
	return []map[string]string{
		{"id": "1", "name": "pedro", "password": "12345", "mail": "pedro@test.com"},
		{"id": "2", "name": "maria", "password": "222"},
	}
}

var (
	testData = map[string]struct {
		form_values url.Values
		expected    string
	}{
		"sesión correcta": {url.Values{
			"mail":     {"pedro@test.com"},
			"password": {"12345"},
		}, "200 OK"},
		// "sesión incorrecta": {url.Values{
		// 	"mail":     {"pedro@test.com"},
		// 	"password": {"222"},
		// }, "406 Not Acceptable"},
	}
)

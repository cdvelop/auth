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

var (
	testData = map[string]struct {
		form_values url.Values
		expected    string
	}{
		"sesión correcta": {url.Values{
			"mail":     {"pedro@test.com"},
			"password": {"1234"},
		}, "200 OK"},
	}
)

func TestLanLoginDefault(t *testing.T) {
	var provider_name = "lan"
	mux := http.NewServeMux()

	mux.HandleFunc("/", db{}.Home)

	auth.Add("/", db{}, false, mux, []string{provider_name})

	server := httptest.NewServer(mux)
	defer server.Close()

	for prueba, data := range testData {
		t.Run((prueba), func(t *testing.T) {

			resp, err := http.PostForm(server.URL+"/login/"+provider_name, data.form_values)

			if err != nil {
				log.Fatal(err)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			if resp.Status != fmt.Sprint(data.expected) {
				fmt.Println("RESPUESTA: ", string(bodyBytes))
				fmt.Println("CÓDIGO: ", resp.Status)
				log.Fatalln()
			}

		})
	}

}

func TestDB(t *testing.T) {

	result := db{}.ReadObject("table", map[string]string{"id": "1"})

	fmt.Println(result)
}

type db struct{}

func (db) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func (d db) ReadObject(table_name string, where_fields map[string]string) map[string]string {
	// Variables para almacenar el resultado y el número de coincidencias encontradas
	var result map[string]string
	var count int

	// Iterar sobre los objetos en la base de datos simulada
	for _, obj := range getObjectsFromDB() {
		match := true
		// Verificar si los campos y valores del objeto coinciden con los especificados en where_fields
		for field, value := range where_fields {
			if obj[field] != value {
				match = false
				break
			}
		}

		// Si se encontró una coincidencia, almacenar el objeto y aumentar el contador
		if match {
			result = obj
			count++
		}
	}

	// Si se encontró exactamente una coincidencia, retornar el resultado
	if count == 1 {
		return result
	}

	return nil
}

// Función auxiliar para simular la obtención de objetos de la base de datos
func getObjectsFromDB() []map[string]string {
	return []map[string]string{
		{"id": "1", "name": "pedro", "password": "111"},
		{"id": "2", "name": "maria", "password": "222"},
	}
}

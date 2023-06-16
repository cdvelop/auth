package auth

import (
	"fmt"
	"net/http"
)

type lan struct {
	*auth
}

// proveedor red area local
func (lan) Name() string {
	return "lan"
}

// Lan login
func (l lan) Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println("Lan Login Recorre los parámetros enviados en la URL")
	for key, values := range r.Form {
		fmt.Printf("%s: %s\n", key, values)
	}
	fmt.Println("---------------------------")

}

func (lan) Callback(w http.ResponseWriter, r *http.Request) {

}

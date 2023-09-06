package auth

import (
	"net/http"

	"github.com/cdvelop/model"
)

type Auth struct {
	https bool
	mux   *http.ServeMux
	db
	client_table_name string

	fields []model.Field
	// retention map comenzará la retención dado el período establecido
	rm map[string]otp

	secret_key_token_signing string

	redirect_success string
	domain           string
	token            string //token name
	status           string //status name

	object model.Object
}

// redirect_success ej "/home", "/", "/login"
// db adapter: ReadObject(table_name, where_fields map[string]string) map[string]string
// client_table_name ej: "users","clients"
// opciones de providers de autenticación en ui login: lan, google_dev
// fields campos opcionales por defecto mail y password
func Add(redirect_success, client_table_name string, data_base db, https bool, http_mux *http.ServeMux, fields ...model.Field) *Auth {

	a := Auth{
		https:             https,
		mux:               http_mux,
		db:                data_base,
		client_table_name: client_table_name,

		fields:                   fields,
		rm:                       make(map[string]otp),
		secret_key_token_signing: buildUniqueKey(16),

		redirect_success: redirect_success,
		token:            "token",
		status:           "status",
	}

	a.buildObject()

	return &a

}

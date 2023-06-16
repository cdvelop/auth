package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/cdvelop/model"
)

type auth struct {
	https bool
	db
	providers []provider
	fields    []model.Field
	// retention map comenzará la retención dado el período establecido
	rm  map[string]otp
	ctx context.Context

	secret_key_token_signing string

	home   string //home redirect
	token  string //token name
	status string //status name
}

// home_redirect ej "/home", "/", "/login"
// db adapter: ReadObject(table_name, where_fields map[string]string) map[string]string
// providers de autenticación requeridos: lan, google_dev
// fields campos opcionales por defecto mail y password
func Add(home_redirect string, data_base db, https bool, mux *http.ServeMux, providers []string, fields ...model.Field) *auth {
	if len(providers) == 0 {
		showErrorAndExit("Debe Agregar mínimo un proveedor de autentificación ej: lan, google_dev")
	}

	a := auth{
		https:                    https,
		db:                       data_base,
		providers:                []provider{},
		fields:                   fields,
		rm:                       make(map[string]otp),
		ctx:                      context.Background(),
		secret_key_token_signing: buildUniqueKey(16),
		home:                     home_redirect,
		token:                    "token",
		status:                   "status",
	}

	for _, opt := range providers {
		switch opt {
		case "lan":
			a.providers = append(a.providers, &lan{auth: &a})
		case "google_dev":
			a.providers = append(a.providers, &google_dev{
				auth:   &a,
				Config: ConfigGoogleOauth2(),
			})
		default:
			showErrorAndExit("proveedor de autentificación: " + opt + " no soportado")
		}

	}

	if mux != nil {
		a.addHttpMuxHandlers(mux)
	}

	// tiempo que se dará en segundos para la autentificación
	go a.retention(30 * time.Second)

	return &a

}

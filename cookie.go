package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cdvelop/token"
)

func (a *Auth) createStatusCookie(w http.ResponseWriter) string {

	value := token.BuildUniqueKey(16)

	SetCookie(a.status, value, a.domain, a.https, 1*time.Minute, w)

	return value

}

// domain ej: "example.com"
func SetCookie(name, value, domain string, https bool, expire time.Duration, w http.ResponseWriter) {

	expiration := time.Now().Add(expire)

	cookie := http.Cookie{
		Name:       name,
		Value:      value,
		Path:       "/",    // / La cookie se aplica a todo el sitio
		Domain:     domain, //"example.com"  La cookie se aplica solo a ese dominio
		Expires:    expiration,
		RawExpires: "",                      // No se utiliza un valor personalizado en bruto para la fecha de vencimiento
		MaxAge:     0,                       // 0 La cookie se eliminará al cerrar el navegador
		Secure:     https,                   //true La cookie solo se enviará a través de una conexión segura (HTTPS)
		HttpOnly:   true,                    //true La cookie no está disponible para scripts del lado del cliente (XSS)
		SameSite:   http.SameSiteStrictMode, // 3 Restricción estricta en el envío de la cookie en solicitudes cruzadas (CSRF)
		Raw:        "",                      // ej: rawCookie := name + "=" + value + "; Path=/; Expires=Wed, 15 Jun 2023 12:00:00 GMT; Secure; HttpOnly"
		Unparsed:   []string{},              //solo recurrir a Unparsed cuando sea absolutamente necesario ej: "SameSite=None", "Priority=High"
	}

	http.SetCookie(w, &cookie)

}

func GetCookie(cookie_name string, r *http.Request) (*http.Cookie, error) {

	cookie, err := r.Cookie(cookie_name)
	if err != nil {
		return nil, fmt.Errorf(err.Error() + " " + cookie_name)
	}

	err = cookie.Valid()
	if err != nil {
		return nil, err
	}

	return cookie, nil
}
